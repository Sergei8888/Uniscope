import { EntityBadRequestException } from '@/general/exceptions/entityBadRequest.exception';

export class OpticalSystemBadRequestException extends EntityBadRequestException {
    constructor(field: string) {
        super('Оптическая система', field.toLowerCase());
    }
}
