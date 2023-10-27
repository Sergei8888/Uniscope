import { EntityBadRequestException } from '@/general/exceptions/entityBadRequest.exception';

export class TelescopeBadRequestException extends EntityBadRequestException {
    constructor(field: string) {
        super('Телескоп', field.toLowerCase());
    }
}
