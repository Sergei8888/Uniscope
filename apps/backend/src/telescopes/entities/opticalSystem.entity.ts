import { Exclude } from 'class-transformer';
import {
    PrimaryGeneratedColumn,
    BaseEntity,
    JoinColumn,
    OneToMany,
    Column,
    Entity,
} from 'typeorm';

import { OpticalSystemType } from '@uniscope/shared/nest';

import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';

@Entity()
export class OpticalSystemEntity extends BaseEntity {
    @PrimaryGeneratedColumn()
    public id: number;

    @Column()
    public aperture: number;

    @Column()
    public manufacturer: string;

    @Column()
    public model: string;

    @Column({ type: 'enum', enum: OpticalSystemType })
    public type: OpticalSystemType;

    @Exclude({ toPlainOnly: true })
    @OneToMany(() => TelescopeEntity, (telescope) => telescope.opticalSystem, {
        cascade: false,
        eager: false,
    })
    @JoinColumn()
    public telescopes: TelescopeEntity[];

    @Column()
    public luminosity: number;

    @Column({
        name: 'focal_length',
    })
    public focalLength: number;
}
