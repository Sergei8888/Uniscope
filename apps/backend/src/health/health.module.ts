import { TerminusModule } from '@nestjs/terminus';
import { Module } from '@nestjs/common';

import { HealthController } from '@/health/health.controller';

@Module({
    imports: [TerminusModule],
    controllers: [HealthController],
})
export class HealthModule {}
