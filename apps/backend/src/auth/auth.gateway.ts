import { ConnectedSocket, OnGatewayConnection } from '@nestjs/websockets';
import { Logger } from '@nestjs/common';
import { RemoteSocket } from 'socket.io';

import { DefaultEventsMap } from 'socket.io/dist/typed-events';

import { JwtService } from '@nestjs/jwt';

import { AuthClientType } from '@uniscope/shared/nest';

import { AuthTokenPayload } from './auth.token.service';

export abstract class AuthGateway implements OnGatewayConnection {
    protected abstract readonly jwtService: JwtService;
    protected abstract readonly logger: Logger;

    // Adds client type to socket data
    public handleConnection(@ConnectedSocket() client: AuthGatewaySocket) {
        let token;

        if (!client.handshake.auth?.access_token) {
            client.emit('exception', [
                'Unauthorized',
                'No access token provided',
            ]);
            client.disconnect();
            this.logger.log(`Client disconnected: No access token provided`);
            return;
        }

        try {
            token = this.jwtService.verify<AuthTokenPayload>(
                client.handshake.auth?.access_token
            );
        } catch (e) {
            client.emit('exception', ['Unauthorized', 'Invalid access token']);
            client.disconnect();
            this.logger.log(`Client disconnected: Access token invalid`);
            return;
        }

        client.data.clientType = token.clientType;
        client.data.userId = token.sub;
    }
}

export type AuthGatewaySocket = RemoteSocket<
    DefaultEventsMap,
    {
        userId: number;
        clientType: AuthClientType;
    }
>;
