import * as Joi from 'joi';

export const getJoiSchema = () =>
    Joi.object({
        NODE_ENV: Joi.string()
            .valid('development', 'production', 'test')
            .default('development'),
        CORS_LIST: Joi.string().default('[http://localhost:3000]'),

        BACKEND_PORT: Joi.number().default(3000),

        VITE_HOST: Joi.string().default('localhost'),
        VITE_PORT: Joi.number().default(4000),

        ADMIN_EMAIL: Joi.string()
            .email({ tlds: { allow: false } })
            .default('admin@example.com'),
        ADMIN_PASSWORD: Joi.string().default('password'),

        POSTGRES_HOST: Joi.string(),
        POSTGRES_USERNAME: Joi.string(),
        POSTGRES_PASSWORD: Joi.string(),
        POSTGRES_PORT: Joi.string().default(5432),

        MAIL_SENDER_LOGIN: Joi.string().email({ tlds: { allow: false } }),
        MAIL_SENDER_PASSWORD: Joi.string(),
        MAIL_RECEIVER_LOGIN: Joi.string().email(),
        MAIL_HOST: Joi.string(),

        PUBLIC_AUTH_JWT_KEY_PATH: Joi.string(),
        PRIVATE_AUTH_JWT_KEY_PATH: Joi.string(),
    });
