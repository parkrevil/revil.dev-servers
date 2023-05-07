import { IsEmail, IsString } from 'class-validator';
import { CreateWithUsernameParams } from 'protobufs/user';

class CreateDto {
  @IsString()
  nickname: string;
}

export class CreateWithUsernameDto extends CreateDto implements CreateWithUsernameParams {
  @IsEmail()
  username: string;

  @IsString()
  password: string;
}
