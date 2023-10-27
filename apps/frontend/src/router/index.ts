// Pages
import { setupLayouts } from 'virtual:generated-layouts';
import { createRouter, createWebHistory } from 'vue-router';

import { watch } from 'vue';

import LandingPage from '@/pages/LandingPage.vue';
import SigninPage from '@/pages/SigninPage.vue';
import SignupPage from '@/pages/SignupPage.vue';
import TelescopeControlPage from '@/pages/TelescopeControlPage.vue';
import TelescopeCatalogPage from '@/pages/TelescopeCatalogPage.vue';
import GalleryPage from '@/pages/GalleryPage.vue';
import { useAuthStore } from '@/auth/model/auth.store';
import ContactsPage from '@/pages/ContactsPage.vue';

// From vite-plugin-vue-layouts
// eslint-disable-next-line import/no-unresolved

const pageDescriptions = {
    home: 'Uniscope – система дистанционного управления астрономическим оборудованием, позволяющая проводить удалённые астрономические наблюдений и съёмки',
};

const routes = setupLayouts([
    {
        path: '/',
        component: LandingPage,
        name: 'home',
        meta: {
            metaTags: [
                {
                    name: 'description',
                    content: pageDescriptions.home,
                },
            ],
        },
    },
    { path: '/signin', component: SigninPage, name: 'signin' },
    { path: '/signup', component: SignupPage, name: 'signup' },
    { path: '/contacts', component: ContactsPage, name: 'contacts' },
    {
        path: '/telescope-catalog',
        component: TelescopeCatalogPage,
        name: 'telescope-catalog',
        beforeEnter: authGuard,
    },
    {
        path: '/telescope-control',
        component: TelescopeControlPage,
        name: 'telescope-control',
        beforeEnter: authGuard,
    },
    {
        path: '/gallery',
        component: GalleryPage,
        name: 'gallery',
        beforeEnter: authGuard,
    },
]);

const router = createRouter({
    history: createWebHistory(),
    scrollBehavior: function (to) {
        if (to.hash) {
            return { el: to.hash, behavior: 'smooth' };
        } else {
            return { left: 0, top: 0 };
        }
    },
    routes,
});

function authGuard(): Promise<void | Record<string, any>> {
    return new Promise((resolve) => {
        if (
            useAuthStore().primaryAuthFinished &&
            !useAuthStore().isAuthenticated
        ) {
            resolve({
                name: 'signin',
            });
            return;
        }

        if (
            useAuthStore().primaryAuthFinished &&
            useAuthStore().isAuthenticated
        ) {
            resolve();
            return;
        }

        watch(
            () => useAuthStore().primaryAuthFinished,
            () => {
                if (!useAuthStore().isAuthenticated) {
                    resolve({
                        name: 'signin',
                    });
                    return;
                }

                resolve();
                return;
            }
        );
    });
}

export { router };
