import { UserId } from 'protobufs/user';

export class CreateUserDto implements UserId {
  id: string;
}
