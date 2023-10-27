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
export class UserEntity extends BaseEntity {
    @PrimaryGeneratedColumn()
    public readonly id: number;

    @Exclude({ toPlainOnly: true })
    @Column({ unique: true })
    public login: string;

    @Column()
    public nickname: string;

    @Exclude({ toPlainOnly: true })
    @Column()
    public password: string;

    @Exclude({ toPlainOnly: true })
    @Column('text', {
        default: [],
        array: true,
        name: 'refresh_tokens',
    })
    public refreshTokens: string[];

    @Exclude({ toPlainOnly: true })
    @OneToMany(() => TelescopeEntity, (telescope) => telescope.owner, {
        cascade: true,
    })
    @JoinColumn()
    public telescopes: TelescopeEntity[];
}
