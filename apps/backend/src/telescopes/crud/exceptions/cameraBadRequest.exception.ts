import { EntityBadRequestException } from '@/general/exceptions/entityBadRequest.exception';

export class CameraBadRequestException extends EntityBadRequestException {
    constructor(field: string) {
        super('Камера', field.toLowerCase());
    }
}
