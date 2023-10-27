import { ConflictException } from '@nestjs/common';

export class UserConflictException extends ConflictException {
    constructor() {
        super(['Пользователь с таким логином уже существует']);
    }
}
