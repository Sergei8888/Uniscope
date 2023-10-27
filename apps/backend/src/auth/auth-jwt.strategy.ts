import {
    UnauthorizedException,
    Injectable,
    Inject,
    Logger,
} from '@nestjs/common';
import { ExtractJwt, Strategy } from 'passport-jwt';
import { PassportStrategy } from '@nestjs/passport';
import { ConfigType } from '@nestjs/config';

import { Repository } from 'typeorm';

import { InjectRepository } from '@nestjs/typeorm';

import { AuthTokenPayload } from './auth.token.service';

import { jwtConfiguration } from '@/config/configurations/jwt.configuration';
import { UserEntity } from '@/users/entities/user.entity';

@Injectable()
export class AuthJwtStrategy extends PassportStrategy(Strategy, 'jwt') {
    private readonly logger = new Logger(AuthJwtStrategy.name);

    constructor(
        @InjectRepository(UserEntity)
        private readonly userRepository: Repository<UserEntity>,
        @Inject(jwtConfiguration.KEY)
        private readonly jwtConfig: ConfigType<typeof jwtConfiguration>
    ) {
        super({
            jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
            ignoreExpiration: false,
            secretOrKey: jwtConfig.privateKey,
        });
    }

    public async validate(payload: AuthTokenPayload): Promise<UserEntity> {
        if (payload.type !== 'access') {
            this.logger.warn('Invalid token type');
            throw new UnauthorizedException();
        }

        return await this.userRepository.findOneByOrFail({
            id: payload.sub,
        });
    }
}
