import { Body, Controller, Post } from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';

import { CreateUserDto } from './dtos';
import { UsernameAlreadyExistsException } from './exceptions';
import { UserService } from './user.service';

@ApiTags('사용자')
@Controller('users')
export class UserController {
  constructor(private userService: UserService) {}

  @Post()
  @ApiOperation({ summary: '사용자 생성' })
  async create(@Body() body: CreateUserDto): Promise<void> {
    if (await this.userService.hasUsername(body.username)) {
      throw new UsernameAlreadyExistsException();
    }

    await this.userService.create(body);
  }
}
