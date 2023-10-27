import { defineStore } from 'pinia';
import { ref } from 'vue';
import { UserDtoI } from '@uniscope/shared/vite';

import { ApiClient } from '@/shared/lib/auth/api-client';

export const useUserStore = defineStore('user', () => {
    const api = ApiClient.getInstance();
    const user = ref<UserDtoI | null>(null);

    async function loadSelf() {
        user.value = await api.getSelf();
    }

    function unloadSelf() {
        user.value = null;
    }

    return { user, loadSelf, unloadSelf };
});
