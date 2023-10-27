import { ApiException } from '@nanogiants/nestjs-swagger-api-exception-decorator';
import {
    Controller,
    HttpCode,
    Body,
    Post,
    UnauthorizedException,
    Query,
} from '@nestjs/common';
import {
    ApiCreatedResponse,
    ApiOkResponse,
    ApiOperation,
    ApiTags,
} from '@nestjs/swagger';

import { IncorrectCredentialsException } from '@/auth/exceptions/incorrect-credentials.exception';
import { AuthService } from '@/auth/auth.service';
import { UserConflictException } from '@/auth/exceptions/user-conflict.exception';
import {
    AuthClientTypeDto,
    AuthSignInDto,
    AuthSignUpDto,
    AuthTokensDto,
    RefreshTokenDto,
} from '@/auth/dto/auth.dto';

@ApiTags('auth')
@Controller('auth/')
export class AuthController {
    constructor(private readonly authService: AuthService) {}

    @ApiOperation({
        summary: 'Sign in',
    })
    @ApiOkResponse({
        description: 'Returns auth tokens pair',
        type: AuthTokensDto,
    })
    @ApiException(() => [IncorrectCredentialsException], {
        description: 'Credentials are invalid',
    })
    @HttpCode(200)
    @Post('signin/')
    public signIn(
        @Body() authSignInDto: AuthSignInDto,
        @Query() clientTypeDto: AuthClientTypeDto
    ): Promise<AuthTokensDto> {
        return this.authService.signIn(authSignInDto, clientTypeDto.clientType);
    }

    @ApiOperation({
        summary: 'Sign up',
    })
    @ApiCreatedResponse({
        description: 'Creates data',
    })
    @ApiException(() => [UserConflictException], {
        description: 'Conflict when creating user',
    })
    @Post('signup/')
    @HttpCode(201)
    public signUp(@Body() authSignUpDto: AuthSignUpDto): Promise<void> {
        return this.authService.signUp(authSignUpDto);
    }

    @ApiOperation({
        summary: 'Renew auth tokens',
    })
    @ApiOkResponse({
        description: 'Returns new auth tokens pair',
        type: AuthTokensDto,
    })
    @ApiException(() => [new UnauthorizedException('Refresh token is invalid')])
    @Post('refresh/')
    @HttpCode(200)
    public refreshTokens(
        @Body() refreshTokenDto: RefreshTokenDto
    ): Promise<AuthTokensDto> {
        return this.authService.refreshTokens(refreshTokenDto.refreshToken);
    }
}
