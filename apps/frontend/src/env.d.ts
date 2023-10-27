interface ImportMetaEnv {
    readonly VITE_HOST: string;
    readonly VITE_PORT: string;

    readonly VITE_API_URL: string;
    readonly VITE_TELESCOPE_CONNECTOR_URL: string;
    readonly VITE_TELESCOPE_CONNECTOR_PATH: string;
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}
