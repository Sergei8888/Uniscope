import { Module } from '@nestjs/common';

import { TypeOrmModule } from '@nestjs/typeorm';

import { TelescopeConnectorGateway } from '@/telescopes/connector/telescope-connector.gateway';
import { ConfiguredJwtModule } from '@/jwt/configured-jwt.module';
import { TelescopeConnectorService } from '@/telescopes/connector/telescope-connector.service';
import { TelescopeAccessService } from '@/telescopes/connector/telescope-access.service';
import { TelescopeCrudModule } from '@/telescopes/crud/telescope-crud.module';
import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';

@Module({
    imports: [
        ConfiguredJwtModule,
        TelescopeCrudModule,
        TypeOrmModule.forFeature([TelescopeEntity]),
    ],
    providers: [
        TelescopeConnectorGateway,
        TelescopeConnectorService,
        TelescopeAccessService,
    ],
})
export class TelescopeConnectorModule {}
