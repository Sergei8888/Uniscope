import { defineStore } from 'pinia';
import { ref } from 'vue';

import { AuthTokensI } from '@uniscope/shared/vite';

import { UnauthorizedApiClient } from '@/shared/lib/auth/unauthorized-api-client';
import { getUnauthorizedAxios } from '@/shared/lib/auth/api-axios';

export const useTokenStore = defineStore(
    'tokens',
    () => {
        const api = new UnauthorizedApiClient(getUnauthorizedAxios());

        const accessToken = ref('');
        const refreshToken = ref('');

        async function refreshTokens() {
            let tokens: AuthTokensI;
            try {
                tokens = await api.refreshTokens({
                    refreshToken: refreshToken.value,
                });
            } catch (error) {
                tokens = {
                    accessToken: '',
                    refreshToken: '',
                };
            }

            assignTokens(tokens);
        }

        function assignTokens(tokens: AuthTokensI) {
            accessToken.value = tokens.accessToken;
            refreshToken.value = tokens.refreshToken;
        }

        function dropTokens() {
            accessToken.value = '';
            refreshToken.value = '';
        }

        return {
            accessToken,
            refreshToken,
            assignTokens,
            refreshTokens,
            dropTokens,
        };
    },
    {
        persist: [
            {
                paths: ['accessToken', 'refreshToken'],
            },
        ],
    }
);
