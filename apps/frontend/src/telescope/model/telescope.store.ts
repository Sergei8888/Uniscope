import { defineStore } from 'pinia';
import { TelescopeDtoI } from '@uniscope/shared/vite';
import { Ref, ref } from 'vue';

import { ApiClient } from '@/shared/lib/auth/api-client';

export const useTelescopeStore = defineStore('telescope', () => {
    const api = ApiClient.getInstance();
    const telescopes = ref<TelescopeDtoI[]>([]) as Ref<TelescopeDtoI[]>;

    async function fetchTelescopes() {
        telescopes.value = await api.getTelescopes();
    }

    return { telescopes, fetchTelescopes };
});
