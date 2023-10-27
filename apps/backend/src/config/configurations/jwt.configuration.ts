import { readFileSync } from 'fs';

import { registerAs } from '@nestjs/config';

type JwtConfigurationType = {
    publicKey: Buffer;
    privateKey: Buffer;
};

export const jwtConfiguration = registerAs(
    'jwt',
    () =>
        ({
            publicKey: readFileSync(
                process.env.PUBLIC_AUTH_JWT_KEY_PATH as string
            ),
            privateKey: readFileSync(
                process.env.PRIVATE_AUTH_JWT_KEY_PATH as string
            ),
        } as JwtConfigurationType)
);
