import { ConfigService, ConfigType } from '@nestjs/config';
import { Inject, Injectable, Logger } from '@nestjs/common';

import { mailConfiguration } from '@/config/configurations/mail.configuration';
import { TelescopeLeadDto } from '@/telescopes/leads/telescope-lead.dto';
import { MailService } from '@/mail/mail.service';

@Injectable()
export class TelescopeLeadsService {
    private readonly logger = new Logger(TelescopeLeadsService.name);

    constructor(
        @Inject(MailService)
        private readonly mailService: MailService,
        @Inject(ConfigService)
        private readonly configService: ConfigService,
        @Inject(mailConfiguration.KEY)
        private readonly mailConfig: ConfigType<typeof mailConfiguration>
    ) {}

    public async uploadLead(
        file: Express.Multer.File,
        telescopeLeadDto: TelescopeLeadDto
    ): Promise<void> {
        await this.mailService.sendMail(
            this.mailConfig.receiver,
            this.mailConfig.sender.login,
            'Новая заявка на добавление телескопа в систему',
            './lead',
            {
                file: {
                    base64: file.buffer.toString('base64'),
                    mime: file.mimetype,
                },
                ...telescopeLeadDto,
            }
        );
        this.logger.log('Telescope lead uploaded');
    }
}
