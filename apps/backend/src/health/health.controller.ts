import { Controller, Get } from '@nestjs/common';
import {
    TypeOrmHealthIndicator,
    MemoryHealthIndicator,
    DiskHealthIndicator,
    HealthCheckService,
    HealthCheck,
} from '@nestjs/terminus';
import { ApiOperation, ApiTags } from '@nestjs/swagger';

@ApiTags('health')
@Controller('/health/')
export class HealthController {
    constructor(
        private readonly health: HealthCheckService,
        private readonly database: TypeOrmHealthIndicator,
        private readonly disk: DiskHealthIndicator,
        private readonly memory: MemoryHealthIndicator
    ) {}

    @ApiOperation({
        summary: 'Docker health check',
    })
    @Get('/')
    @HealthCheck()
    public healthCheck() {
        return this.health.check([
            () => this.database.pingCheck('database'),
            () =>
                this.disk.checkStorage('storage', {
                    path: '/',
                    thresholdPercent: 0.9,
                }),
            () => this.memory.checkHeap('memory_heap', 512 * Math.pow(1024, 2)),
            () => this.memory.checkRSS('memory_rss', 400 * Math.pow(1024, 2)),
        ]);
    }
}
