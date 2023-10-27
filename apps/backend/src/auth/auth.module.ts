import { Module } from '@nestjs/common';

import { TypeOrmModule } from '@nestjs/typeorm';

import { AuthController } from '@/auth/auth.controller';
import { AuthService } from '@/auth/auth.service';
import { AuthJwtStrategy } from '@/auth/auth-jwt.strategy';
import { ConfiguredJwtModule } from '@/jwt/configured-jwt.module';
import { UserEntity } from '@/users/entities/user.entity';
import { AuthTokenService } from '@/auth/auth.token.service';

@Module({
    imports: [ConfiguredJwtModule, TypeOrmModule.forFeature([UserEntity])],
    providers: [AuthService, AuthTokenService, AuthJwtStrategy],
    controllers: [AuthController],
    exports: [AuthService],
})
export class AuthModule {}
