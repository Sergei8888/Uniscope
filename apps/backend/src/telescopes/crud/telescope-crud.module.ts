import { Module } from '@nestjs/common';

import { TypeOrmModule } from '@nestjs/typeorm';

import { TelescopeCrudController } from '@/telescopes/crud/telescope-crud.controller';
import { OpticalSystemEntity } from '@/telescopes/entities/opticalSystem.entity';
import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';
import { CameraEntity } from '@/telescopes/entities/camera.entity';
import { MountEntity } from '@/telescopes/entities/mount.entity';

@Module({
    imports: [
        TypeOrmModule.forFeature([
            TelescopeEntity,
            OpticalSystemEntity,
            CameraEntity,
            MountEntity,
        ]),
    ],
    controllers: [TelescopeCrudController],
})
export class TelescopeCrudModule {}
