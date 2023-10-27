import { registerAs } from '@nestjs/config';

type MailConfigurationType = {
    sender: {
        login: string;
        password: string;
    };
    receiver: string;
    host: string;
};

export const mailConfiguration = registerAs(
    'mail',
    () =>
        ({
            sender: {
                login: process.env.MAIL_SENDER_LOGIN,
                password: process.env.MAIL_SENDER_PASSWORD,
            },
            receiver: process.env.MAIL_RECEIVER_LOGIN,
            host: process.env.MAIL_HOST,
        } as MailConfigurationType)
);
