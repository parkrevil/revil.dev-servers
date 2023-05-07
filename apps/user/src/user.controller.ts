import { Metadata } from '@grpc/grpc-js';
import { Body, Controller, UsePipes, ValidationPipe } from '@nestjs/common';
import { UserId, UserServiceController, UserServiceControllerMethods } from 'protobufs/user';

import { CreateWithUsernameDto } from './dtos';
import { UserService } from './user.service';

@Controller()
@UserServiceControllerMethods()
export class UserController implements UserServiceController {
  constructor(private readonly userService: UserService) {}

  async createWithUsername(@Body() req: CreateWithUsernameDto, metadata?: Metadata): Promise<UserId> {
    try {
      console.log(req, metadata);

      const user = await this.userService.createWithUsername(req);

      console.log(user);

      return {
        id: 'test',
      };
    } catch (e) {
      console.error(e);
    }
  }
}
