import { registerAs } from '@nestjs/config';

type DatabaseConfigurationType = {
    host: string;
    port: number;
    username: string;
    password: string;
    database: string;
};

export const databaseConfiguration = registerAs(
    'database',
    () =>
        ({
            host: process.env.POSTGRES_HOST,
            port: Number(process.env.POSTGRES_PORT),
            username: process.env.POSTGRES_USERNAME,
            password: process.env.POSTGRES_PASSWORD,
            database: process.env.POSTGRES_DATABASE,
        } as DatabaseConfigurationType)
);
