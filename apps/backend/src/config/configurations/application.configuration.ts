import { registerAs } from '@nestjs/config';

export interface ApplicationConfigurationI {
    mode: string;
    port: number;
    corsList: string[];
}

export const applicationConfiguration = registerAs('application', () => {
    return {
        mode: process.env.NODE_ENV,
        port: Number(process.env.BACKEND_PORT),
        corsList: JSON.parse(process.env.CORS_LIST as string),
    } as ApplicationConfigurationI;
});
