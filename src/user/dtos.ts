import { ApiProperty } from '@nestjs/swagger';
import { IsEmail, IsOptional, IsString, IsUrl } from 'class-validator';

export class CreateUserDto {
  @ApiProperty({
    format: 'email',
    description: '아이디',
  })
  @IsEmail()
  username: string;

  @ApiProperty({
    description: '비밀번호',
  })
  @IsString()
  password: string;

  @ApiProperty({
    description: '닉네임',
  })
  nickname: string;
}
