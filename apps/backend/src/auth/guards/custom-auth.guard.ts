import { Injectable, Logger, UnauthorizedException } from '@nestjs/common';
import { AuthGuard } from '@nestjs/passport';

@Injectable()
export class CustomAuthGuard extends AuthGuard('jwt') {
    private readonly logger = new Logger('CustomAuthGuard');

    constructor() {
        super();
    }

    public handleRequest(err, user, info) {
        if (err) {
            throw new UnauthorizedException(err);
        }

        if (info) {
            throw new UnauthorizedException(info.message);
        }

        if (!user) {
            this.logger.warn({
                message: 'No user for this token was found',
            });
            throw new UnauthorizedException('No user for this token was found');
        }

        return user;
    }
}
