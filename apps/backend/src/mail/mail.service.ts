import { MailerService } from '@nestjs-modules/mailer';
import { Inject, Injectable, Logger } from '@nestjs/common';

@Injectable()
export class MailService {
    private readonly logger = new Logger(MailerService.name);

    constructor(
        @Inject(MailerService)
        private readonly mailerService: MailerService
    ) {}

    public async sendMail(
        receiver: string,
        sender: string,
        subject: string,
        template: string,
        context: Record<string, any>
    ) {
        await this.mailerService.sendMail({
            to: receiver,
            from: sender,
            subject,
            template,
            context,
        });
        this.logger.debug('Mail sent');
    }
}
