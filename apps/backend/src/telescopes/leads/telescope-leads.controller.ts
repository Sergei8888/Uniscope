import { FileInterceptor } from '@nestjs/platform-express';
import { Express } from 'express';
import {
    UseInterceptors,
    UploadedFile,
    Controller,
    UseGuards,
    Body,
    Post,
    UnauthorizedException,
} from '@nestjs/common';

import { AuthGuard } from '@nestjs/passport';

import { ApiOperation, ApiTags } from '@nestjs/swagger';

import { ApiException } from '@nanogiants/nestjs-swagger-api-exception-decorator';

import { TelescopeLeadsService } from '@/telescopes/leads/telescope-leads.service';
import { TelescopeLeadDto } from '@/telescopes/leads/telescope-lead.dto';

@ApiException(() => UnauthorizedException)
@ApiTags('telescopes')
@UseGuards(AuthGuard('jwt'))
@Controller('telescopes/leads')
export class TelescopeLeadsController {
    constructor(
        private readonly telescopeLeadsService: TelescopeLeadsService
    ) {}

    @ApiOperation({
        summary: 'Request to create a new telescope in the system',
    })
    @Post()
    @UseInterceptors(FileInterceptor('file'))
    public uploadApplication(
        @UploadedFile() file: Express.Multer.File,
        @Body() telescopeLeadDto: TelescopeLeadDto
    ) {
        return this.telescopeLeadsService.uploadLead(file, telescopeLeadDto);
    }
}
