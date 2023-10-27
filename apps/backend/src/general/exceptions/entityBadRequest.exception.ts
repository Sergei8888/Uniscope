import { BadRequestException } from '@nestjs/common';

export class EntityBadRequestException extends BadRequestException {
    constructor(entityName: string, fieldName: string) {
        super([`${entityName} с указанным ${fieldName} не найден`]);
    }
}
