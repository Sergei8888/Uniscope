import { Exclude } from 'class-transformer';
import {
    PrimaryGeneratedColumn,
    BaseEntity,
    JoinColumn,
    OneToMany,
    Column,
    Entity,
} from 'typeorm';

import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';

@Entity()
export class CameraEntity extends BaseEntity {
    @PrimaryGeneratedColumn()
    public id: number;

    @Column()
    public manufacturer: string;

    @Column()
    public model: string;

    @Column({ name: 'sensor_size' })
    public sensorSize: number;

    @Column({ name: 'pixel_size' })
    public pixelSize: number;

    @Column({ name: 'quantum_efficiency' })
    public quantumEfficiency: number;

    @Column()
    public width: number;

    @Column()
    public height: number;

    @Exclude({ toPlainOnly: true })
    @OneToMany(() => TelescopeEntity, (telescope) => telescope.camera, {
        cascade: false,
        eager: false,
    })
    @JoinColumn()
    private telescopes: TelescopeEntity[];
}
