import { AuthClientType } from '@uniscope/shared/nest';

import { JwtService } from '@nestjs/jwt';

import { Injectable, UnauthorizedException } from '@nestjs/common';

import { AuthTokensDto } from './dto/auth.dto';

import { UserEntity } from '@/users/entities/user.entity';

@Injectable()
export class AuthTokenService {
    constructor(private readonly jwtService: JwtService) {}

    public async generateNewTokensPair(
        user: UserEntity,
        clientType: AuthClientType
    ): Promise<AuthTokensDto> {
        return {
            accessToken: this.generateToken(
                {
                    type: 'access',
                    sub: user.id,
                    clientType,
                },
                '1h'
            ),
            refreshToken: this.generateToken(
                {
                    type: 'refresh',
                    sub: user.id,
                    clientType,
                },
                '30d'
            ),
        };
    }

    private generateToken(payload: AuthTokenPayload, expiresIn: string) {
        return this.jwtService.sign(
            <AuthTokenPayload>{
                ...payload,
            },
            {
                expiresIn,
            }
        );
    }

    public getValidRefreshTokens(user: UserEntity) {
        return user.refreshTokens.filter((token) => {
            try {
                this.validateToken(token);
                return true;
            } catch (e) {
                return false;
            }
        });
    }

    public validateToken(token: string): AuthTokenPayload {
        let payload: AuthTokenPayload;
        try {
            payload = this.jwtService.verify<AuthTokenPayload>(token);
        } catch (e) {
            throw new UnauthorizedException('Refresh token is invalid');
        }

        return payload;
    }
}

export type AuthTokenPayload = {
    type: 'access' | 'refresh';
    sub: number;
    clientType: AuthClientType;
};
