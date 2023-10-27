import { UserDtoI } from '@uniscope/shared/nest';

export class UserDto implements UserDtoI {
    public readonly id: number;
    public nickname: string;
}
