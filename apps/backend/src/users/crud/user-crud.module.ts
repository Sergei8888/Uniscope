import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';

import { UserCrudController } from '@/users/crud/user-crud.controller';

import { UserEntity } from '@/users/entities/user.entity';
import { AuthJwtStrategy } from '@/auth/auth-jwt.strategy';

@Module({
    imports: [TypeOrmModule.forFeature([UserEntity])],
    providers: [AuthJwtStrategy],
    controllers: [UserCrudController],
})
export class UserCrudModule {}
