import { ConfigModule, ConfigType } from '@nestjs/config';
import { MulterModule } from '@nestjs/platform-express';
import { TypeOrmModule } from '@nestjs/typeorm';
import { MiddlewareConsumer, Module, NestModule } from '@nestjs/common';
import { ClsModule } from 'nestjs-cls';

import { v4 as uuidv4 } from 'uuid';

import { PrometheusModule } from '@willsoto/nestjs-prometheus';

import { databaseConfiguration } from '@/config/configurations/database.configuration';
import { OpticalSystemEntity } from '@/telescopes/entities/opticalSystem.entity';
import { mailConfiguration } from '@/config/configurations/mail.configuration';
import { jwtConfiguration } from '@/config/configurations/jwt.configuration';
import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';
import { CameraEntity } from '@/telescopes/entities/camera.entity';
import { MountEntity } from '@/telescopes/entities/mount.entity';
import { getJoiSchema } from '@/config/config.schema';
import { UserEntity } from '@/users/entities/user.entity';
import { AuthModule } from '@/auth/auth.module';
import { MailModule } from '@/mail/mail.module';
import { HealthModule } from '@/health/health.module';
import { HttpLoggerMiddleware } from '@/logger/http-logger.middleware';
import { TelescopesModule } from '@/telescopes/telescopes.module';
import { UserModule } from '@/users/user.module';
import { applicationConfiguration } from '@/config/configurations/application.configuration';
import { docsConfiguration } from '@/config/configurations/docs.configuration';
import { AsyncapiModule } from '@/docs/asyncapi/asyncapi.module';

@Module({
    imports: [
        ConfigModule.forRoot({
            isGlobal: true,
            validationSchema: getJoiSchema(),
            validationOptions: { allowUnknows: false, abortEarly: false },
            load: [
                applicationConfiguration,
                databaseConfiguration,
                mailConfiguration,
                jwtConfiguration,
                docsConfiguration,
            ],
            ignoreEnvFile: true,
            cache: true,
        }),
        ClsModule.forRoot({
            middleware: {
                mount: true,
                generateId: true,
                idGenerator: (req: Request) => {
                    return req.headers['X-Request-Id'] ?? uuidv4();
                },
            },
        }),
        PrometheusModule.register({
            path: '/metrics',
        }),
        TypeOrmModule.forRootAsync({
            imports: [ConfigModule.forFeature(databaseConfiguration)],
            inject: [databaseConfiguration.KEY],
            useFactory: (
                databaseConfig: ConfigType<typeof databaseConfiguration>
            ) => ({
                type: 'postgres',
                ...databaseConfig,
                extra: {
                    ssl: false,
                },
                ssl: {
                    rejectUnauthorized: false,
                },
                entities: [
                    OpticalSystemEntity,
                    TelescopeEntity,
                    CameraEntity,
                    MountEntity,
                    UserEntity,
                ],
                autoLoadEntities: true,
                synchronize: true,
            }),
        }),
        AsyncapiModule,
        MulterModule.register(),
        UserModule,
        AuthModule,
        MailModule,
        HealthModule,
        TelescopesModule,
    ],
})
export class AppModule implements NestModule {
    public configure(consumer: MiddlewareConsumer) {
        consumer.apply(HttpLoggerMiddleware).forRoutes('*');
    }
}
