import { IsEmail, IsString } from 'class-validator';
import { CreateUserWithUsernameParams } from 'protobufs/user';

class CreateUserDto {
  @IsString()
  nickname: string;
}

export class CreateUserWithUsernameDto extends CreateUserDto implements CreateUserWithUsernameParams {
  @IsEmail()
  username: string;

  @IsString()
  password: string;
}
