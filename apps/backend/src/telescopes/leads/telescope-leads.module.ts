import { Module } from '@nestjs/common';

import { TelescopeLeadsController } from '@/telescopes/leads/telescope-leads.controller';
import { TelescopeLeadsService } from '@/telescopes/leads/telescope-leads.service';
import { MailModule } from '@/mail/mail.module';

@Module({
    imports: [MailModule],
    providers: [TelescopeLeadsService],
    controllers: [TelescopeLeadsController],
})
export class TelescopeLeadsModule {}
