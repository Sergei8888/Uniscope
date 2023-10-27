import { Logger as TypeOrmLogger } from 'typeorm';
import { Logger as NestLogger } from '@nestjs/common';

export class TypeormLogger implements TypeOrmLogger {
    private readonly logger = new NestLogger('TypeORM');

    public logQuery(query: string, parameters?: unknown[]) {
        this.log(
            'log',
            `${query} -- Parameters: ${this.stringifyParameters(parameters)}`
        );
    }

    public logQueryError(error: string, query: string, parameters?: unknown[]) {
        this.logger.error(
            `${query} -- Parameters: ${this.stringifyParameters(
                parameters
            )} -- ${error}`
        );
    }

    public logQuerySlow(time: number, query: string, parameters?: unknown[]) {
        this.log(
            'warn',
            `Time: ${time} -- Parameters: ${this.stringifyParameters(
                parameters
            )} -- ${query}`
        );
    }

    public logMigration(message: string) {
        this.log('log', message);
    }

    public logSchemaBuild(message: string) {
        this.log('log', message);
    }

    public log(level: 'log' | 'info' | 'warn', message: string) {
        if (level === 'log') {
            return this.logger.log(message);
        }
        if (level === 'info') {
            return this.logger.debug(message);
        }
        if (level === 'warn') {
            return this.logger.warn(message);
        }
    }

    private stringifyParameters(parameters?: unknown[]) {
        try {
            return JSON.stringify(parameters);
        } catch {
            return '';
        }
    }
}
