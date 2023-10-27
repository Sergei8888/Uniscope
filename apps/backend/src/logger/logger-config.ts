import { WinstonModuleOptions } from 'nest-winston';
import { utilities as nestWinstonModuleUtilities } from 'nest-winston';
import * as winston from 'winston';
import { ClsService } from 'nestjs-cls';

const { printf } = winston.format;

export function loggerConfig(cls: ClsService): WinstonModuleOptions {
    const removeValueFieldFormat = printf((info) => {
        delete info.value;
        return JSON.stringify(info, null, 2);
    });

    const addReqIdFieldFormat = printf((info) => {
        info.reqId = cls.getId();
        return JSON.stringify(info, null, 2);
    });

    const filterMetadataFormat = printf((info) => {
        if (info?.body?.password) {
            info.body.password = '***';
        }

        if (info?.body?.passwordRepeated) {
            info.body.passwordRepeated = '***';
        }

        return JSON.stringify(info, null, 2);
    });

    return {
        level: 'debug',
        transports: [
            new winston.transports.Console({
                format: winston.format.combine(
                    winston.format.timestamp(),
                    removeValueFieldFormat,
                    addReqIdFieldFormat,
                    filterMetadataFormat,
                    nestWinstonModuleUtilities.format.nestLike('Uniscope', {
                        colors: true,
                        prettyPrint: true,
                    })
                ),
            }),
        ],
    };
}
