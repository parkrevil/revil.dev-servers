import { Body, Controller, Get, Post } from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';
import { CreateUserDto } from './dtos';
import { UserService } from './user.service';
import { EmailAlreadyExistsException } from './exceptions';

@ApiTags('사용자')
@Controller('users')
export class UserController {
  constructor(
    private userService: UserService,
  ) {}

  @Post()
  @ApiOperation({ summary: '사용자 생성' })
  async create(@Body() body: CreateUserDto): Promise<void> {
    if (await this.userService.hasEmail(body.email)) {
      throw new EmailAlreadyExistsException();
    }

    await this.userService.create(body);
  }
}
