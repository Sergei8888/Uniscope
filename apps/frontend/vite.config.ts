import vue from '@vitejs/plugin-vue';
import { defineConfig, loadEnv } from 'vite';
import Layouts from 'vite-plugin-vue-layouts';
import Unfonts from 'unplugin-fonts/vite';
import svgLoader from 'vite-svg-loader';
import AutoImport from 'unplugin-auto-import/vite';
import Components from 'unplugin-vue-components/vite';
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers';

// https://vitejs.dev/config/
// Because of vite spec
// eslint-disable-next-line import/no-default-export
export default defineConfig(({ mode }) => {
    const envDir = '../../';
    const env = loadEnv(mode, envDir, 'VITE');

    return {
        envDir,

        server: {
            host: env.VITE_HOST,
            port: Number(env.VITE_PORT),
        },

        plugins: [
            // Basic vue plugin
            vue(),
            // Vite plugin layouts (we are not using pages from that author)
            Layouts({
                layoutsDirs: 'src/layouts',
                defaultLayout: 'DefaultLayout',
            }),
            // Element plus installation
            AutoImport({
                resolvers: [ElementPlusResolver()],
            }),
            Components({
                resolvers: [ElementPlusResolver()],
            }),
            // Loading svg as components and optimize through svgo
            svgLoader({
                defaultImport: 'component',
            }),
            // Add fonts from Google fonts (or other online providers)
            Unfonts({
                google: {
                    families: [
                        {
                            name: 'Montserrat',
                            styles: 'wght@400;500;600',
                        },
                    ],
                },
            }),
        ],

        resolve: {
            alias: {
                '@': '/src',
            },
        },
    };
});
