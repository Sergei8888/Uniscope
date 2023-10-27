import { Module } from '@nestjs/common';

import { UserCrudModule } from '@/users/crud/user-crud.module';

@Module({
    imports: [UserCrudModule],
})
export class UserModule {}
