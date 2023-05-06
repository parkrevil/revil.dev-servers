import { grpcClientFactories, GrpcServer } from '@app/core';
import { Module } from '@nestjs/common';
import { ClientsModule } from '@nestjs/microservices';

import { AuthResolver } from './auth.resolver';

@Module({
  imports: [ClientsModule.registerAsync([grpcClientFactories[GrpcServer.Auth]])],
  providers: [AuthResolver],
})
export class AuthModule {}
