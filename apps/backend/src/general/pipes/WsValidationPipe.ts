import { Injectable, ValidationPipe } from '@nestjs/common';
import { WsException } from '@nestjs/websockets';

@Injectable()
export class WsValidationPipe extends ValidationPipe {
    public createExceptionFactory() {
        return (validationErrors = []) => {
            if (this.isDetailedOutputDisabled) {
                return new WsException(
                    'Bad request: ' + JSON.stringify(validationErrors)
                );
            }
            const errors = this.flattenValidationErrors(validationErrors);

            return new WsException(errors);
        };
    }
}
