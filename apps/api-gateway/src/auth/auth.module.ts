import { AuthConfig } from '@app/core/configs';
import { Module } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';

import { AuthResolver } from './auth.resolver';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
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
    ]),
  ],
  providers: [AuthResolver],
})
export class AuthModule {}
