import { UnauthorizedException } from '@nestjs/common';

export class IncorrectCredentialsException extends UnauthorizedException {
    constructor() {
        super(['Неверный логин или пароль']);
    }
}
