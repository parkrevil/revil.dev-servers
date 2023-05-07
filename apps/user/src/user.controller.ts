import { Metadata } from '@grpc/grpc-js';
import { Controller, Get } from '@nestjs/common';
import {
  createUserWithUsernameParams,
  UserId,
  UserServiceController,
  UserServiceControllerMethods,
} from 'protobufs/user';

import { UserService } from './user.service';

@Controller()
@UserServiceControllerMethods()
export class UserController implements UserServiceController {
  constructor(private readonly userService: UserService) {}

  createUserWithUsername(request: createUserWithUsernameParams, metadata?: Metadata): UserId {
    return {
      id: 'test',
    };
  }
}
