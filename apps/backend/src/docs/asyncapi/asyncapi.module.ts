import { join } from 'path';

import { Module } from '@nestjs/common';
import { ServeStaticModule } from '@nestjs/serve-static';
import { ConfigModule, ConfigType } from '@nestjs/config';

import { docsConfiguration } from '@/config/configurations/docs.configuration';

@Module({
    imports: [
        ServeStaticModule.forRootAsync({
            imports: [ConfigModule.forFeature(docsConfiguration)],
            inject: [docsConfiguration.KEY],
            useFactory: (docsConfig: ConfigType<typeof docsConfiguration>) => {
                return [
                    {
                        serveRoot: docsConfig.asyncApiUrl,
                        rootPath: join(__dirname, 'dist'),
                    },
                ];
            },
        }),
    ],
})
export class AsyncapiModule {}
