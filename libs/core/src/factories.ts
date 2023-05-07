import { ConfigService } from '@nestjs/config';
import { ClientsProviderAsyncOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

import { GrpcServer } from './enums';
import { AuthConfig, UserConfig } from './types';

export const grpcClientInjectionTokens: { [key in GrpcServer]: string } = {
  [GrpcServer.Auth]: 'AUTH_GRPC_PACKAGE',
  [GrpcServer.User]: 'USER_GRPC_PACKAGE',
};

export const grpcClientFactories: { [key in GrpcServer]: ClientsProviderAsyncOptions } = {
  [GrpcServer.Auth]: {
    name: grpcClientInjectionTokens.auth,
    useFactory: (configService: ConfigService) => ({
      transport: Transport.GRPC,
      options: {
        url: configService.get<AuthConfig>('auth').url,
        package: 'auth',
        protoPath: join(process.cwd(), 'protobufs/auth.proto'),
      },
    }),
    inject: [ConfigService],
  },
  [GrpcServer.User]: {
    name: grpcClientInjectionTokens.user,
    useFactory: (configService: ConfigService) => ({
      transport: Transport.GRPC,
      options: {
        url: configService.get<UserConfig>('user').url,
        package: 'user',
        protoPath: join(process.cwd(), 'protobufs/user.proto'),
      },
    }),
    inject: [ConfigService],
  },
};
