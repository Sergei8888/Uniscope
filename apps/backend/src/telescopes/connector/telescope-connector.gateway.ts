import {
    ConnectedSocket,
    MessageBody,
    OnGatewayConnection,
    OnGatewayDisconnect,
    SubscribeMessage,
    WebSocketGateway,
    WebSocketServer,
} from '@nestjs/websockets';
import { Logger, UsePipes } from '@nestjs/common';
import { Socket, RemoteSocket, Server } from 'socket.io';
import { DefaultEventsMap } from 'socket.io/dist/typed-events';
import { JwtService } from '@nestjs/jwt';
import { IsArray } from 'class-validator';

import { AuthClientType } from '@uniscope/shared/nest';

import {
    RtcIceCandidateDto,
    RtcSessionDescriptionDto,
} from '@/telescopes/connector/dto/rtc.dto';
import { WsValidationPipe } from '@/general/pipes/WsValidationPipe';
import { AuthGateway, AuthGatewaySocket } from '@/auth/auth.gateway';
import { TelescopeConnectorService } from '@/telescopes/connector/telescope-connector.service';

@UsePipes(new WsValidationPipe())
@WebSocketGateway(getServerConfig().port, {
    namespace: getServerConfig().namespace,
    cors: getServerConfig().cors,
    path: getServerConfig().path,
})
export class TelescopeConnectorGateway
    extends AuthGateway
    implements OnGatewayConnection, OnGatewayDisconnect
{
    @WebSocketServer()
    protected readonly server: TelescopeConnectorGatewayServer;
    protected readonly logger = new Logger(TelescopeConnectorGateway.name);

    protected readonly jwtService: JwtService;

    constructor(
        jwtService: JwtService,
        private readonly telescopeConnectorService: TelescopeConnectorService
    ) {
        super();
        this.jwtService = jwtService;
    }

    public async handleConnection(
        @ConnectedSocket() client: TelescopeConnectorGatewaySocket
    ): Promise<any> {
        // Handle auth
        super.handleConnection(client);

        if (!client.handshake.auth.telescopeId) {
            client.emit('exception', ['Telescope ID is not provided']);
            client.disconnect();
        }

        switch (client.data.clientType) {
            case AuthClientType.frontend:
                await this.telescopeConnectorService.handleFrontendConnection(
                    client
                );
                break;
            case AuthClientType.hardend:
                await this.telescopeConnectorService.handleHardendConnection(
                    client
                );
                break;
            default:
                this.logger.error({
                    message: `Unknown client type`,
                    clientType: client.data.clientType,
                });
                client.emit('exception', ['Unknown client type']);
                client.disconnect();
                break;
        }

        client.join(client.handshake.auth.telescopeId);
    }

    @SubscribeMessage('message')
    public async handleMessagePub(
        @ConnectedSocket() client: Socket,
        @MessageBody() payload: MessageDto
    ) {
        let roomSockets = await this.server
            .in(client.handshake.auth.telescopeId)
            .fetchSockets();

        let retryCount = 0;
        const sendMessageInterval = setInterval(async () => {
            retryCount++;
            if (retryCount > 10) {
                client.emit('exception', [
                    'Could not find another client to connect to',
                ]);
                client.disconnect();
                clearInterval(sendMessageInterval);
            }

            roomSockets = await this.server
                .in(client.handshake.auth.telescopeId)
                .fetchSockets();

            if (roomSockets.length === 2) {
                this.sendMessageToTheOtherClient(client, roomSockets, payload);
                clearInterval(sendMessageInterval);
            }
        }, 1000);
    }

    private sendMessageToTheOtherClient(
        client: Socket,
        room: Array<TelescopeConnectorGatewaySocket>,
        description: MessageDto
    ) {
        for (const currentSocket of room) {
            if (currentSocket.data.clientType !== client.data.clientType) {
                currentSocket.emit('message', description);
            }
        }
    }

    public async handleDisconnect(client: TelescopeConnectorGatewaySocket) {
        switch (client.data.clientType) {
            case AuthClientType.frontend:
                await this.telescopeConnectorService.handleFrontendDisconnection(
                    client
                );
                break;
            case AuthClientType.hardend:
                await this.telescopeConnectorService.handleHardendDisconnection(
                    client
                );
                break;
        }
    }
}

export type TelescopeConnectorGatewaySocket = RemoteSocket<
    {
        exception: (_: Array<string>) => void;
        message: (_: MessageDto) => void;
    },
    any
> &
    AuthGatewaySocket;

type TelescopeConnectorGatewayServer = Server<
    DefaultEventsMap,
    {
        exception: (_: Array<string>) => void;
        message: (_: MessageDto) => void;
    }
>;

type MessageDto =
    | { description: RtcSessionDescriptionDto }
    | { ice_candidate: RtcIceCandidateDto };

function getServerConfig(): {
    port: number;
    namespace: string;
    cors: Array<string>;
    path: string;
} {
    if (!process.env.TELESCOPE_CONNECTOR_PORT) {
        throw new Error('Missing TELESCOPE_CONNECTOR_PORT');
    }

    if (!process.env.TELESCOPE_CONNECTOR_NAMESPACE) {
        throw new Error('Missing TELESCOPE_CONNECTOR_NAMESPACE');
    }

    if (!process.env.CORS_LIST) {
        throw new Error('Missing CORS_LIST');
    }

    if (!IsArray(JSON.parse(process.env.CORS_LIST))) {
        throw new Error('CORS_LIST is not an array');
    }

    if (!process.env.TELESCOPE_CONNECTOR_PATH) {
        throw new Error('Missing TELESCOPE_CONNECTOR_PATH');
    }

    return {
        port: Number(process.env.TELESCOPE_CONNECTOR_PORT),
        namespace: process.env.TELESCOPE_CONNECTOR_NAMESPACE,
        cors: JSON.parse(process.env.CORS_LIST as string),
        path: process.env.TELESCOPE_CONNECTOR_PATH,
    };
}
