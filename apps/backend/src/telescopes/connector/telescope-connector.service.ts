import { Injectable, Logger } from '@nestjs/common';

import { Repository } from 'typeorm';

import { InjectRepository } from '@nestjs/typeorm';

import { TelescopeConnectorGatewaySocket } from '@/telescopes/connector/telescope-connector.gateway';
import { TelescopeAccessService } from '@/telescopes/connector/telescope-access.service';
import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';

@Injectable()
export class TelescopeConnectorService {
    private readonly logger = new Logger(TelescopeConnectorService.name);

    constructor(
        private readonly telescopeAccessService: TelescopeAccessService,
        @InjectRepository(TelescopeEntity)
        private readonly telescopesRepository: Repository<TelescopeEntity>
    ) {}

    public async handleFrontendConnection(
        client: TelescopeConnectorGatewaySocket
    ) {
        this.logger.log({
            message: `Frontend connected`,
            telescopeId: client.handshake.auth.telescopeId,
        });

        try {
            await this.telescopeAccessService.bookTelescope(
                client.handshake.auth.telescopeId
            );
        } catch (e) {
            this.logger.error({
                message: `Telescope could not be booked`,
                telescopeId: client.handshake.auth.telescopeId,
                error: e,
            });
            client.emit('exception', ['Telescope could not be booked']);
            client.disconnect();
        }
    }

    public async handleFrontendDisconnection(
        client: TelescopeConnectorGatewaySocket
    ) {
        this.logger.log({
            message: `Frontend disconnected`,
            telescopeId: client.handshake.auth.telescopeId,
        });

        try {
            await this.telescopeAccessService.freeTelescope(
                client.handshake.auth.telescopeId
            );
        } catch (e) {
            this.logger.error({
                message: `Telescope could not be freed`,
                telescopeId: client.handshake.auth.telescopeId,
                error: e,
            });
        }
    }

    public async handleHardendConnection(
        client: TelescopeConnectorGatewaySocket
    ) {
        this.logger.log({
            message: `Hardend connected`,
            telescopeId: client.handshake.auth.telescopeId,
        });

        try {
            const telescope = await this.telescopesRepository.findOneOrFail({
                relations: ['owner'],
                where: {
                    id: client.handshake.auth.telescopeId,
                },
            });

            if (telescope.owner.id !== client.data.userId) {
                throw new Error('Telescope does not belong to user');
            }
        } catch (e) {
            client.emit('exception', [e.message]);
            client.disconnect();
            return;
        }

        try {
            await this.telescopeAccessService.makeTelescopeOnline(
                client.handshake.auth.telescopeId
            );
        } catch (e) {
            this.logger.error({
                message: `Telescope could not be made online`,
                telescopeId: client.handshake.auth.telescopeId,
                error: e,
            });
            client.emit('exception', ['Telescope could not be made online']);
            client.disconnect();
        }
    }

    public async handleHardendDisconnection(
        client: TelescopeConnectorGatewaySocket
    ) {
        this.logger.log({
            message: `Hardend disconnected`,
            telescopeId: client.handshake.auth.telescopeId,
        });

        await this.telescopeAccessService.makeTelescopeOffline(
            client.handshake.auth.telescopeId
        );
    }
}
