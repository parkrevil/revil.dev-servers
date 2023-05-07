import { Metadata } from '@grpc/grpc-js';
import { Controller } from '@nestjs/common';
import { UserId, UserServiceController, UserServiceControllerMethods } from 'protobufs/user';

import { CreateUserWithUsernameDto } from './dtos';
import { UserService } from './user.service';

@Controller()
@UserServiceControllerMethods()
export class UserController implements UserServiceController {
  constructor(private readonly userService: UserService) {}

  async createUserWithUsername(req: CreateUserWithUsernameDto, metadata?: Metadata): Promise<UserId> {
    const user = await this.userService.createWithUsername(req);

    console.log(req, metadata, user);

    return {
      id: 'test',
    };
  }
}
