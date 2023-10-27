import { Injectable, NestMiddleware, Logger } from '@nestjs/common';

import { Request, Response, NextFunction } from 'express';

@Injectable()
export class HttpLoggerMiddleware implements NestMiddleware {
    private readonly logger = new Logger(HttpLoggerMiddleware.name);

    public use(request: Request, response: Response, next: NextFunction): void {
        const ignorePaths = ['/api/health', '/metrics'];
        const { method, originalUrl } = request;

        if (ignorePaths.includes(originalUrl)) {
            next();
            return;
        }

        const start = performance.now();
        response.on('close', () => {
            const duration = Math.floor(performance.now() - start);
            const { statusCode } = response;

            if (statusCode >= 400) {
                this.logger.warn(
                    `${method} ${originalUrl} ${statusCode} ${duration}ms`
                );
            } else {
                this.logger.log(
                    `${method} ${originalUrl} ${statusCode} ${duration}ms`
                );
            }

            this.logger.debug({
                message: 'Request',
                body: request.body,
                headers: request.headers,
            });
        });

        next();
    }
}
