import { ApiException } from '@nanogiants/nestjs-swagger-api-exception-decorator';
import { ParseIntPipe, Controller, Param, Get } from '@nestjs/common';
import { ApiOkResponse, ApiTags, ApiOperation } from '@nestjs/swagger';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';

import { User } from '@/general/decorators/user.decorator';
import { UserDto } from '@/users/dto/user.dto';
import { UserEntity } from '@/users/entities/user.entity';
import { UserNotFoundByIdException } from '@/users/crud/exceptions/user-not-found-by-id.exception';
import { AuthGuarded } from '@/auth/decorators/auth-guarded.decorator';

@ApiTags('users')
@AuthGuarded()
@Controller('/users')
export class UserCrudController {
    constructor(
        @InjectRepository(UserEntity)
        private readonly usersRepository: Repository<UserEntity>
    ) {}

    @ApiOperation({
        summary: 'Get self user information',
    })
    @ApiOkResponse({
        description: 'Returns user information',
        type: UserDto,
    })
    @Get('/me')
    public getCurrentUserInformation(@User() user: UserEntity): UserDto {
        return user;
    }

    @ApiOperation({
        summary: 'Get other user information',
    })
    @ApiOkResponse({
        description: 'Returns user information',
        type: UserDto,
    })
    @ApiException(() => new UserNotFoundByIdException())
    @Get('/:id')
    public async getOtherUserInformation(
        @Param('id', ParseIntPipe) userId: number
    ) {
        const user = await this.usersRepository.findOneBy({
            id: userId,
        });

        if (!user) {
            throw new UserNotFoundByIdException();
        }

        return user;
    }
}
