import { grpcClientInjectionTokens } from '@app/core';
import { Inject, OnModuleInit } from '@nestjs/common';
import { Args, Mutation, Resolver } from '@nestjs/graphql';
import { ClientGrpc } from '@nestjs/microservices';
import { UserServiceClient } from 'protobufs/user';
import { lastValueFrom } from 'rxjs';

import { CreateUserWithUsernameInput } from './dtos';
import { User } from './models';

@Resolver((of) => User)
export class UserResolver implements OnModuleInit {
  private userServiceClient: UserServiceClient;

  constructor(@Inject(grpcClientInjectionTokens.user) private userGrpcClient: ClientGrpc) {}

  onModuleInit(): void {
    this.userServiceClient = this.userGrpcClient.getService<UserServiceClient>('UserService');
  }

  @Mutation(() => Boolean)
  async createUserWithUsername(@Args('input') input: CreateUserWithUsernameInput): Promise<void> {
    const result = await lastValueFrom(this.userServiceClient.createWithUsername(input));

    console.log(result);

    return;
  }
}
