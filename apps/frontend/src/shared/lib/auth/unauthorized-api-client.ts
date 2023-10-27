import {
    AuthSignInDtoI,
    AuthTokensI,
    RefreshTokenDtoI,
} from '@uniscope/shared/vite';

import { UnauthorizedAxiosInstance } from '@/shared/lib/auth/api-axios';

export interface UnauthorizedApiClientI {
    signIn(signInDto: AuthSignInDtoI): Promise<AuthTokensI>;

    refreshTokens(refreshTokenDto: RefreshTokenDtoI): Promise<AuthTokensI>;

    signUp(signUpDto: AuthSignInDtoI): Promise<void>;
}

export class UnauthorizedApiClient implements UnauthorizedApiClientI {
    constructor(private readonly axios: UnauthorizedAxiosInstance) {}

    public async signIn(signInDto: AuthSignInDtoI) {
        const response = await this.axios.post<AuthTokensI>(
            '/auth/signin?clientType=frontend',
            signInDto
        );

        return response.data;
    }

    public async refreshTokens(refreshTokenDto: RefreshTokenDtoI) {
        const response = await this.axios.post<AuthTokensI>(
            '/auth/refresh',
            refreshTokenDto
        );

        return response.data;
    }

    public async signUp(signUpDto: AuthSignInDtoI) {
        const response = await this.axios.post<void>('/auth/signup', signUpDto);

        return response.data;
    }
}
