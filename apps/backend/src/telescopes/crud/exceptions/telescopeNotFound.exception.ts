import { EntityNotFoundException } from '@/general/exceptions/entityNotFound.exception';

export class TelescopeNotFoundException extends EntityNotFoundException {
    constructor(field: string) {
        super('Телескоп', field.toLowerCase());
    }
}
