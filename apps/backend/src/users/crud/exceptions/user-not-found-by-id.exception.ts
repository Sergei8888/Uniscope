import { EntityNotFoundException } from '@/general/exceptions/entityNotFound.exception';

export class UserNotFoundByIdException extends EntityNotFoundException {
    constructor() {
        super('Пользователь', 'id');
    }
}
