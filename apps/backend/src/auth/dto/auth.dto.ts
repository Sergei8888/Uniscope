import { IsString, Matches, MaxLength, MinLength } from 'class-validator';

import {
    AuthClientType,
    AuthClientTypeDtoI,
    AuthExceptions,
    AuthSignInDtoI,
    AuthSignUpDtoI,
    AuthTokensI,
    RefreshTokenDtoI,
} from '@uniscope/shared/nest';

import { Match } from '@/general/decorators/match.decorator';

export class AuthSignInDto implements AuthSignInDtoI {
    /*
     * @example "TestUser"
     */
    login: string;
    /*
     * @example "Testpassword888"
     */
    password: string;
}

export class AuthClientTypeDto implements AuthClientTypeDtoI {
    clientType: AuthClientType;
}

export class AuthSignUpDto implements AuthSignUpDtoI {
    @IsString() nickname: string;

    @MinLength(6)
    @MaxLength(20)
    login: string;

    @MinLength(6)
    @MaxLength(20)
    @Matches(/((?=.*\d)|(?=.*\W+))(?![.\n])(?=.*[A-Z])(?=.*[a-z]).*$/, {
        message: AuthExceptions.passwordIsToWeak,
    })
    password: string;

    @Match('password', { message: AuthExceptions.passwordsDoNotMatch })
    passwordRepeated: string;
}

export class AuthTokensDto implements AuthTokensI {
    accessToken: string;
    refreshToken: string;
}

export class RefreshTokenDto implements RefreshTokenDtoI {
    refreshToken: string;
}
