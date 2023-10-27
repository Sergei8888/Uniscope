import {
    PrimaryGeneratedColumn,
    JoinColumn,
    BaseEntity,
    ManyToOne,
    Entity,
    Column,
} from 'typeorm';

import { TelescopeStatus, TelescopeTarget } from '@uniscope/shared/nest';

import { OpticalSystemEntity } from '@/telescopes/entities/opticalSystem.entity';
import { CameraEntity } from '@/telescopes/entities/camera.entity';
import { MountEntity } from '@/telescopes/entities/mount.entity';
import { UserEntity } from '@/users/entities/user.entity';

@Entity()
export class TelescopeEntity extends BaseEntity {
    @PrimaryGeneratedColumn()
    public readonly id: number;

    @Column()
    public name: string;

    @Column()
    public description?: string;

    @Column({
        unique: true,
        name: 'img_url',
        default: '/static/telescope-default.png',
    })
    public imgUrl: string;

    @Column({ type: 'float' })
    public latitude: number;

    @Column({ type: 'float' })
    public longitude: number;

    @Column({ nullable: true })
    public zoom: number;

    @Column({ type: 'enum', enum: TelescopeTarget })
    public target: TelescopeTarget;

    @Column({
        type: 'enum',
        enum: TelescopeStatus,
        default: TelescopeStatus.Offline,
    })
    public status: TelescopeStatus;

    @ManyToOne(() => OpticalSystemEntity, { eager: true })
    @JoinColumn({ name: 'optical_system_id' })
    public opticalSystem: OpticalSystemEntity;

    @ManyToOne(() => CameraEntity, { eager: true })
    @JoinColumn({ name: 'camera_id' })
    public camera: CameraEntity;

    @ManyToOne(() => MountEntity, { eager: true })
    @JoinColumn({ name: 'mount_id' })
    public mount: MountEntity;

    @ManyToOne(() => UserEntity, { eager: true })
    @JoinColumn({ name: 'owner_id' })
    public owner: UserEntity;
}
