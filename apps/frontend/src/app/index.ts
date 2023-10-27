import { createApp } from 'vue';
import { createPinia } from 'pinia';
import piniaPluginPersistedState from 'pinia-plugin-persistedstate';
import { initYandexMetrika } from 'yandex-metrika-vue3';

import { router } from '@/router';
import App from '@/app/ui/App.vue';

const app = createApp(App);
const pinia = createPinia();
pinia.use(piniaPluginPersistedState);

app.use(pinia)
    .use(router)
    .use(initYandexMetrika, {
        id: '94279472',
        router: router,
        env:
            import.meta.env.VITE_HOST === 'uniscope.astromodel.ru'
                ? 'production'
                : 'development',
        scriptSrc: 'https://mc.yandex.ru/metrika/tag.js',
    });

export { app };
