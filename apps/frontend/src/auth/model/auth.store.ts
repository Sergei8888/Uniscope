import { defineStore } from 'pinia';
import { ref } from 'vue';
import { AuthSignInDtoI, AuthSignUpDtoI } from '@uniscope/shared/vite';

import { AuthTokensI } from '@uniscope/shared/src';

import { ApiClient } from '@/shared/lib/auth/api-client';
import { useTokenStore } from '@/shared/model/token.store';
import { useUserStore } from '@/user/model/user.store';

export const useAuthStore = defineStore('auth', () => {
    const api = ApiClient.getInstance();

    const isAuthenticated = ref(false);
    const primaryAuthFinished = ref(false);

    async function signIn(credentials: AuthSignInDtoI) {
        const tokens = await api.signIn(credentials);
        await authFromTokens(tokens);
    }

    async function signUp(dto: AuthSignUpDtoI) {
        await api.signUp(dto);
        await signIn(dto);
    }

    async function authFromTokens(tokens: AuthTokensI) {
        useTokenStore().assignTokens(tokens);
        await authByTokens();
    }

    async function authByTokens() {
        try {
            if (useTokenStore().refreshToken) {
                isAuthenticated.value = true;
                await useUserStore().loadSelf();
            }
        } catch (e: any) {
            console.error('Процесс авторизации прошел неудачно', e);
        } finally {
            primaryAuthFinished.value = true;
        }
    }

    function logout() {
        isAuthenticated.value = false;
        useUserStore().unloadSelf();
        useTokenStore().dropTokens();
    }

    return {
        isAuthenticated,
        primaryAuthFinished,
        signIn,
        signUp,
        authByTokens,
        logout,
    };
});
