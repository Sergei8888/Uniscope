import { createParamDecorator, ExecutionContext } from '@nestjs/common';
import { Request } from 'express';

export const Token = createParamDecorator(
    (data, context: ExecutionContext): string => {
        const request: Request = context.switchToHttp().getRequest();
        return ((value) => (value ? (value.split(' ').at(-1) as string) : ''))(
            request.headers.authorization
        );
    }
);
