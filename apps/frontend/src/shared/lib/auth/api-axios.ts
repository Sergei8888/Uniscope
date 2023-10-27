import axios, {
    AxiosInstance,
    HttpStatusCode,
    InternalAxiosRequestConfig,
} from 'axios';

import { useTokenStore } from '@/shared/model/token.store';
import { useAuthStore } from '@/auth/model/auth.store';

export function getUnauthorizedAxios(): UnauthorizedAxiosInstance {
    const api = axios.create({
        baseURL: import.meta.env.VITE_API_URL,
        headers: {
            'Content-Type': 'application/json',
        },
    });

    api.interceptors.response.use(undefined, async (error) => {
        return await interceptRefreshTokenInvalidation(error);
    });

    return api;
}

export function getAuthorizedAxios(): AuthorizedAxiosInstance {
    const api = getUnauthorizedAxios();

    api.interceptors.request.use((config) => {
        return setAuthHeader(config);
    });
    api.interceptors.response.use(undefined, async (error) => {
        return await interceptAccessTokenExpiration(error, api);
    });

    return api;
}

export function setAuthHeader(config: InternalAxiosRequestConfig) {
    config.headers.Authorization = `Bearer ${useTokenStore().accessToken}`;
    return config;
}

export async function interceptAccessTokenExpiration(
    error: any,
    axios: AxiosInstance
) {
    const accessTokenExpired =
        isTokenExpired(error) && error.config.url !== '/auth/refresh';

    if (accessTokenExpired) {
        await useTokenStore().refreshTokens();
        if (useTokenStore().accessToken) {
            return await axios.request(error.config);
        }
    }

    throw error;
}

export async function interceptRefreshTokenInvalidation(error: any) {
    if (error.config.url === '/auth/refresh' && error.response.status === 401) {
        useAuthStore().logout();
    }

    throw error;
}

function isTokenExpired(error: any): boolean {
    return (
        error.response.data.statusCode === HttpStatusCode.Unauthorized &&
        error.response.data.message === 'jwt expired'
    );
}

export type AuthorizedAxiosInstance = AxiosInstance;
export type UnauthorizedAxiosInstance = AxiosInstance;
