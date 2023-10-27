import { join } from 'path';

import { HandlebarsAdapter } from '@nestjs-modules/mailer/dist/adapters/handlebars.adapter';
import { ConfigModule, ConfigType } from '@nestjs/config';
import { MailerModule } from '@nestjs-modules/mailer';
import { Module } from '@nestjs/common';

import { mailConfiguration } from '@/config/configurations/mail.configuration';
import { MailService } from '@/mail/mail.service';

@Module({
    imports: [
        MailerModule.forRootAsync({
            imports: [ConfigModule.forFeature(mailConfiguration)],
            inject: [mailConfiguration.KEY],
            useFactory: (mailConfig: ConfigType<typeof mailConfiguration>) => ({
                transport: {
                    host: mailConfig.host,
                    secure: true,
                    auth: {
                        user: mailConfig.sender.login,
                        pass: mailConfig.sender.password,
                    },
                },
                template: {
                    dir: join(__dirname, 'templates'),
                    adapter: new HandlebarsAdapter(),
                },
                options: {
                    strict: false,
                },
            }),
        }),
    ],
    providers: [MailService],
    exports: [MailService],
})
export class MailModule {}
