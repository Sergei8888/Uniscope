export enum AuthClientType {
    'hardend' = 'hardend',
    'frontend' = 'frontend',
}

export interface AuthSignInDtoI {
    login: string;
    password: string;
}

export interface AuthClientTypeDtoI {
    clientType: AuthClientType;
}

export interface AuthTokensI {
    accessToken: string;
    refreshToken: string;
}

export enum AuthExceptions {
    passwordsDoNotMatch = 'Введенные пароли не совпадают',
    passwordIsToWeak = 'Пароль слишком простой',
}

export interface AuthSignUpDtoI {
    login: string;
    password: string;
    passwordRepeated: string;
    nickname: string;
}

export interface RefreshTokenDtoI {
    refreshToken: string;
}
