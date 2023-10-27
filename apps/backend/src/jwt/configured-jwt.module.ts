import { Module } from '@nestjs/common';

import { ConfigType } from '@nestjs/config';

import { JwtModule } from '@nestjs/jwt';

import { jwtConfiguration } from '@/config/configurations/jwt.configuration';

@Module({
    imports: [
        JwtModule.registerAsync({
            inject: [jwtConfiguration.KEY],
            useFactory: (jwtConfig: ConfigType<typeof jwtConfiguration>) => ({
                publicKey: jwtConfig.publicKey,
                privateKey: jwtConfig.privateKey,
                signOptions: {
                    issuer: 'https://uniscope.astromodel.ru',
                    algorithm: 'RS256',
                },
                verifyOptions: {
                    algorithms: ['RS256'],
                },
            }),
        }),
    ],
    exports: [JwtModule],
})
export class ConfiguredJwtModule {}
