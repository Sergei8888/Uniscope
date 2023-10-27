import { EntityBadRequestException } from '@/general/exceptions/entityBadRequest.exception';

export class MountBadRequestException extends EntityBadRequestException {
    constructor(field: string) {
        super('Монтировка', field.toLowerCase());
    }
}
