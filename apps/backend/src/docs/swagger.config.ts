import { DocumentBuilder } from '@nestjs/swagger';

export const swaggerConfig = new DocumentBuilder()
    .setTitle('Uniscope Backend API')
    .setVersion('1.0')
    .addTag('users')
    .addTag('auth')
    .addTag('gallery')
    .addBearerAuth()
    .build();
