import { Injectable, Logger, UnauthorizedException } from '@nestjs/common';

import * as bcrypt from 'bcryptjs';

import { Repository } from 'typeorm';

import { AuthClientType } from '@uniscope/shared/nest';

import { InjectRepository } from '@nestjs/typeorm';

import { IncorrectCredentialsException } from '@/auth/exceptions/incorrect-credentials.exception';
import { UserEntity } from '@/users/entities/user.entity';
import { UserConflictException } from '@/auth/exceptions/user-conflict.exception';
import {
    AuthSignInDto,
    AuthSignUpDto,
    AuthTokensDto,
} from '@/auth/dto/auth.dto';
import { AuthTokenService } from '@/auth/auth.token.service';

@Injectable()
export class AuthService {
    private readonly logger = new Logger(AuthService.name);

    constructor(
        private readonly tokenService: AuthTokenService,
        @InjectRepository(UserEntity)
        private readonly userRepository: Repository<UserEntity>
    ) {}

    public async refreshTokens(refreshToken: string): Promise<AuthTokensDto> {
        const tokenPayload = this.tokenService.validateToken(refreshToken);

        const user = await this.userRepository
            .createQueryBuilder('user')
            .select()
            .where(':refreshToken = ANY(user.refresh_tokens)', {
                refreshToken,
            })
            .getOne();

        if (!user) {
            throw new UnauthorizedException(
                'User with this token was not found. Please, sign in again'
            );
        }

        const tokens = await this.tokenService.generateNewTokensPair(
            user,
            tokenPayload.clientType
        );

        user.refreshTokens = user.refreshTokens.filter((token) => {
            this.logger.log({
                token: token,
                refreshToken: refreshToken,
            });
            return token !== refreshToken;
        });
        user.refreshTokens = this.tokenService.getValidRefreshTokens(user);
        user.refreshTokens.push(tokens.refreshToken);

        await user.save();

        return tokens;
    }

    public async signUp(userCredentials: AuthSignUpDto) {
        try {
            userCredentials.password = bcrypt.hashSync(
                userCredentials.password,
                bcrypt.genSaltSync()
            );

            const res = await this.userRepository.insert(userCredentials);
            this.logger.log({
                message: 'User has registered',
                id: res.identifiers[0].id,
            });
        } catch (error) {
            throw new UserConflictException();
        }
    }

    public async signIn(
        userCredentials: AuthSignInDto,
        clientType: AuthClientType
    ): Promise<AuthTokensDto> {
        const user = await this.userRepository.findOneBy({
            login: userCredentials.login,
        });
        if (!user) {
            throw new IncorrectCredentialsException();
        }
        await this.validateUserCredentials(user, userCredentials);

        const tokens = await this.tokenService.generateNewTokensPair(
            user,
            clientType
        );

        user.refreshTokens.push(tokens.refreshToken);
        await user.save();

        this.logger.log(`User {id: ${user.id}} is logged in`);
        return tokens;
    }

    private async validateUserCredentials(
        targetUser: UserEntity,
        userCredentials: AuthSignInDto
    ) {
        const isMatch: boolean = bcrypt.compareSync(
            userCredentials.password,
            targetUser.password
        );

        if (!isMatch) {
            throw new IncorrectCredentialsException();
        }

        this.logger.debug(
            `User\`s {id: ${targetUser.id}} credentials validated`
        );
    }
}
