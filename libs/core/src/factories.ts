import { ConfigService } from '@nestjs/config';
import { ClientsProviderAsyncOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

import { AuthConfig } from './configs';
import { GrpcServer } from './enums';

export const grpcClientFactories: { [key in GrpcServer]: ClientsProviderAsyncOptions } = {
  [GrpcServer.Auth]: {
    name: 'AUTH_GRPC_PACKAGE',
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
};
