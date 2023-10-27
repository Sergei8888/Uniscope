import { NotFoundException } from '@nestjs/common';

export class EntityNotFoundException extends NotFoundException {
    constructor(entityName: string, fieldName: string) {
        super([`${entityName} с указанным ${fieldName} не найден`]);
    }
}
