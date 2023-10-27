import { Module } from '@nestjs/common';

import { TypeOrmModule } from '@nestjs/typeorm';

import { TelescopeCrudModule } from '@/telescopes/crud/telescope-crud.module';
import { TelescopeLeadsModule } from '@/telescopes/leads/telescope-leads.module';
import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';
import { TelescopeConnectorModule } from '@/telescopes/connector/telescope-connector.module';

@Module({
    imports: [
        TypeOrmModule.forFeature([TelescopeEntity]),
        TelescopeCrudModule,
        TelescopeLeadsModule,
        TelescopeConnectorModule,
    ],
})
export class TelescopesModule {}
