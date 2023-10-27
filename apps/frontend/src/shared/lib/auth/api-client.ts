import {
    AuthSignInDtoI,
    AuthTokensI,
    RefreshTokenDtoI,
    TelescopeDtoI,
    UserDtoI,
} from '@uniscope/shared/vite';

import {
    AuthorizedAxiosInstance,
    getAuthorizedAxios,
    getUnauthorizedAxios,
} from '@/shared/lib/auth/api-axios';
import {
    UnauthorizedApiClient,
    UnauthorizedApiClientI,
} from '@/shared/lib/auth/unauthorized-api-client';

export interface ApiClientI extends UnauthorizedApiClientI {
    getTelescopes(): Promise<Array<any>>;
}

export class ApiClient implements ApiClientI {
    private static instance: ApiClient;

    constructor(
        private readonly axios: AuthorizedAxiosInstance,
        private readonly unauthApi: UnauthorizedApiClientI
    ) {}

    // Auth
    public refreshTokens(
        refreshTokenDto: RefreshTokenDtoI
    ): Promise<AuthTokensI> {
        return this.unauthApi.refreshTokens(refreshTokenDto);
    }

    public signIn(signInDto: AuthSignInDtoI): Promise<AuthTokensI> {
        return this.unauthApi.signIn(signInDto);
    }

    public signUp(signUpDto: AuthSignInDtoI): Promise<void> {
        return this.unauthApi.signUp(signUpDto);
    }

    // Users
    public async getSelf(): Promise<UserDtoI> {
        const response = await this.axios.get<UserDtoI>('/users/me');
        return response.data;
    }

    public async getUserById(id: number): Promise<UserDtoI> {
        const response = await this.axios.get<UserDtoI>(`/users/${id}`);
        return response.data;
    }

    // Telescopes
    public async getTelescopes(): Promise<Array<TelescopeDtoI>> {
        const response = await this.axios.get('/telescopes');
        return response.data;
    }

    public static getInstance(): ApiClient {
        if (!this.instance) {
            this.instance = new ApiClient(
                getAuthorizedAxios(),
                new UnauthorizedApiClient(getUnauthorizedAxios())
            );
        }

        return this.instance;
    }
}
