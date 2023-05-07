import { AuthConfig } from '@app/core/configs';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

import { AuthModule } from './auth.module';

async function bootstrap() {
  const appContext = await NestFactory.createApplicationContext(AuthModule);
  const config = appContext.get(ConfigService).get<AuthConfig>('auth');

  await appContext.close();

  const app = await NestFactory.createMicroservice<MicroserviceOptions>(AuthModule, {
    transport: Transport.GRPC,
    options: {
      url: config.url,
      package: 'auth',
      protoPath: join(process.cwd(), 'protobufs/auth.proto'),
    },
  });

  await app.listen();
}

bootstrap();
