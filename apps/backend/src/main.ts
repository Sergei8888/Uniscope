import {
    INestApplication,
    Logger,
    RequestMethod,
    ValidationPipe,
} from '@nestjs/common';
import { SwaggerModule } from '@nestjs/swagger';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';

import helmet from 'helmet';
import { WinstonModule } from 'nest-winston';
import { ClsService } from 'nestjs-cls';

import { TransformInterceptor } from '@/global/interceptors/transform.interceptor';
import { swaggerConfig } from '@/docs/swagger.config';
import { AppModule } from '@/app.module';
import { loggerConfig } from '@/logger/logger-config';
import { ApplicationConfigurationI } from '@/config/configurations/application.configuration';
import { DocsConfiguration } from '@/config/configurations/docs.configuration';

export async function bootstrap() {
    const application = await NestFactory.create(AppModule);
    const configService = application.get(ConfigService);

    const readyApp = await enableDocs(
        enableCors(enableGlobals(enableHelmet(enableLogger(application))))
    );

    await listen(readyApp);
    new Logger('Application').log(
        'Listening on port: ' + configService.get('application.port')
    );
}

function enableLogger(application: INestApplication) {
    const logger = WinstonModule.createLogger(
        loggerConfig(application.get(ClsService))
    );
    application.useLogger(logger);

    return application;
}

function enableHelmet(application: INestApplication) {
    const configService = application.get<ConfigService>(ConfigService);
    const docsConfiguration =
        configService.getOrThrow<DocsConfiguration>('docs');

    const excludedRoutes = [
        docsConfiguration.asyncApiUrl,
        docsConfiguration.swaggerUrl,
    ];

    application.use((req, res, next) => {
        for (const route of excludedRoutes) {
            if (req.url.startsWith(route)) {
                return next();
            }
        }

        return helmet()(req, res, next);
    });
    return application;
}

function enableGlobals(application: INestApplication) {
    application.setGlobalPrefix('api', {
        exclude: [{ path: '/metrics', method: RequestMethod.GET }],
    });
    application.useGlobalInterceptors(new TransformInterceptor());
    application.useGlobalPipes(
        new ValidationPipe({
            transform: true,
        })
    );

    return application;
}

function enableCors(application: INestApplication) {
    const configService = application.get<ConfigService>(ConfigService);
    const appConfig =
        configService.getOrThrow<ApplicationConfigurationI>('application');
    application.enableCors({
        origin: appConfig.corsList,
        credentials: true,
    });

    return application;
}

async function enableDocs(application: INestApplication) {
    const configService = application.get<ConfigService>(ConfigService);
    const docsConfiguration =
        configService.getOrThrow<DocsConfiguration>('docs');

    enableSwagger(docsConfiguration.swaggerUrl, application);

    return application;
}

function enableSwagger(swaggerPath: string, application: INestApplication) {
    const logger = new Logger('Swagger');
    const document = SwaggerModule.createDocument(application, swaggerConfig);

    SwaggerModule.setup(swaggerPath, application, document, {
        useGlobalPrefix: true,
    });

    logger.log(
        'Swagger documentation served successfully on path: ' + swaggerPath
    );
}

function listen(application: INestApplication) {
    const configService = application.get<ConfigService>(ConfigService);

    return application.listen(configService.getOrThrow('application.port'));
}

bootstrap();
