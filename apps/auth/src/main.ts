import { AuthConfig } from '@app/core/configs';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

import { AuthModule } from './auth.module';

async function bootstrap() {
  const appContext = await NestFactory.createApplicationContext(AuthModule);
  const authConfig = appContext.get(ConfigService).get<AuthConfig>('auth');

  await appContext.close();

  const app = await NestFactory.createMicroservice<MicroserviceOptions>(AuthModule, {
    transport: Transport.GRPC,
    options: {
      url: authConfig.url,
      package: 'auth',
      protoPath: join(process.cwd(), 'protobufs/auth.proto'),
    },
  });

  await app.listen();
}

bootstrap();
