import { Controller, Get, Post } from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';

@ApiTags('사용자')
@Controller('users')
export class UserController {
  @Post()
  @ApiOperation({ summary: '사용자 생성' })
  create(): void {}
}
