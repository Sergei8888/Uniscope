import { registerAs } from '@nestjs/config';

export interface DocsConfiguration {
    // /api included
    swaggerUrl: string;
    // /api not included
    asyncApiUrl: string;
}

export const docsConfiguration = registerAs('docs', (): DocsConfiguration => {
    return {
        swaggerUrl: '/docs',
        asyncApiUrl: '/api/async-docs',
    };
});
