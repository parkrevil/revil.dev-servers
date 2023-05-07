import { grpcClientFactories } from '@app/core';
import { Module } from '@nestjs/common';
import { ClientsModule } from '@nestjs/microservices';

import { UserResolver } from './user.resolver';

@Module({
  imports: [ClientsModule.registerAsync([grpcClientFactories.user])],
  providers: [UserResolver],
})
export class UserModule {}
