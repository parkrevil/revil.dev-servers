import { grpcClientFactories } from '@app/core';
import { Module } from '@nestjs/common';
import { ClientsModule } from '@nestjs/microservices';

import { AuthResolver } from './auth.resolver';

@Module({
  imports: [ClientsModule.registerAsync([grpcClientFactories.auth])],
  providers: [AuthResolver],
})
export class AuthModule {}
