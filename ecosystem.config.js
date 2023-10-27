module.exports = {
    apps: [
        {
            name: 'Frontend',
            script: 'cd ./apps/frontend && pnpm run dev',
        },
        {
            name: 'Backend',
            script: 'cd ./apps/backend && pnpm run dev',
        },
        {
            name: 'Postgres',
            script: 'cd ./apps/database && pnpm run dev',
        },
        {
            name: 'Pgadmin',
            script: 'cd ./apps/pgadmin && pnpm run dev',
        },
        {
            name: 'Minio',
            script: 'cd ./apps/minio && pnpm run dev',
        },
        {
            name: 'Shared',
            script: 'cd ./packages/shared && pnpm run dev',
        },
    ],
};
