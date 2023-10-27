import { Exclude } from 'class-transformer';
import {
    PrimaryGeneratedColumn,
    JoinColumn,
    BaseEntity,
    OneToMany,
    Column,
    Entity,
} from 'typeorm';

import { MountType } from '@uniscope/shared/nest';

import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';

@Entity()
export class MountEntity extends BaseEntity {
    @PrimaryGeneratedColumn()
    public readonly id: number;

    @Column()
    public manufacturer: string;

    @Column()
    public model: string;

    @Column()
    public tier: number;

    @Column({ type: 'enum', enum: MountType })
    public type: MountType;

    @Exclude({ toPlainOnly: true })
    @OneToMany(() => TelescopeEntity, (telescope) => telescope.mount, {
        eager: false,
        cascade: false,
    })
    @JoinColumn()
    private telescopes: TelescopeEntity[];
}
