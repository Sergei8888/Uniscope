import { JwtService } from '@nestjs/jwt';
import { Inject, Injectable, Logger } from '@nestjs/common';

import { EntityManager, Repository } from 'typeorm';

import { ConfigType } from '@nestjs/config';

import { InjectRepository } from '@nestjs/typeorm';

import { TelescopeStatus } from '@uniscope/shared/nest';

import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';
import { jwtConfiguration } from '@/config/configurations/jwt.configuration';

@Injectable()
export class TelescopeAccessService {
    private readonly logger = new Logger(TelescopeAccessService.name);

    constructor(
        private readonly jwtService: JwtService,
        @Inject(jwtConfiguration.KEY)
        private readonly jwtConfig: ConfigType<typeof jwtConfiguration>,
        private readonly entityManager: EntityManager,
        @InjectRepository(TelescopeEntity)
        private readonly telescopesRepository: Repository<TelescopeEntity>
    ) {}

    public async bookTelescope(telescopeId: number) {
        this.logger.log({
            message: `Telescope has been booked`,
            telescopeId,
        });

        const telescope = await this.telescopesRepository.findOneByOrFail({
            id: telescopeId,
        });

        if (telescope.status !== TelescopeStatus.Online) {
            throw new Error('Telescope is not online');
        }

        telescope.status = TelescopeStatus.Booked;
        await telescope.save();
    }

    public async freeTelescope(telescopeId: number) {
        const telescope = await this.telescopesRepository.findOneByOrFail({
            id: telescopeId,
        });

        if (telescope.status !== TelescopeStatus.Booked) {
            throw new Error('Telescope is not booked');
        }

        telescope.status = TelescopeStatus.Online;
        await telescope.save();
    }

    public async makeTelescopeOnline(telescopeId: number) {
        const telescope = await this.telescopesRepository.findOneByOrFail({
            id: telescopeId,
        });

        if (telescope.status !== TelescopeStatus.Offline) {
            throw new Error('Telescope is not offline');
        }

        telescope.status = TelescopeStatus.Online;
        await telescope.save();
    }

    public async makeTelescopeOffline(telescopeId: number) {
        await this.telescopesRepository.update(telescopeId, {
            status: TelescopeStatus.Offline,
        });
    }
}
