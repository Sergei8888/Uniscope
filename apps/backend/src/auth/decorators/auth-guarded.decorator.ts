import {
    applyDecorators,
    UnauthorizedException,
    UseGuards,
} from '@nestjs/common';
import { ApiBearerAuth } from '@nestjs/swagger';
import { ApiException } from '@nanogiants/nestjs-swagger-api-exception-decorator';

import { CustomAuthGuard } from '@/auth/guards/custom-auth.guard';

export function AuthGuarded() {
    return applyDecorators(
        ApiBearerAuth(),
        ApiException(() => new UnauthorizedException()),
        UseGuards(CustomAuthGuard)
    );
}
