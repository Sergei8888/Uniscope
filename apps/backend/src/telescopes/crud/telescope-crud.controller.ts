import { ApiOkResponse, ApiOperation, ApiTags } from '@nestjs/swagger';
import { Controller, Get } from '@nestjs/common';

import { Repository } from 'typeorm';

import { InjectRepository } from '@nestjs/typeorm';

import { TelescopeDto } from './dto/telescope.dto';

import { AuthGuarded } from '@/auth/decorators/auth-guarded.decorator';
import { TelescopeEntity } from '@/telescopes/entities/telescope.entity';

@ApiTags('telescopes')
@AuthGuarded()
@Controller('telescopes')
export class TelescopeCrudController {
    constructor(
        @InjectRepository(TelescopeEntity)
        private readonly telescopeRepository: Repository<TelescopeEntity>
    ) {}

    @ApiOperation({
        description: 'Get all telescopes',
    })
    @ApiOkResponse({
        description: 'Returns all telescopes',
        type: () => TelescopeDto,
        isArray: true,
    })
    @Get('/')
    public async getTelescopes(): Promise<Array<TelescopeDto>> {
        return await this.telescopeRepository.find();
    }
}
