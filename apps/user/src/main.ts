import { UserConfig } from '@app/core/configs';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

import { UserModule } from './user.module';

async function bootstrap() {
  const appContext = await NestFactory.createApplicationContext(UserModule);
  const config = appContext.get(ConfigService).get<UserConfig>('user');

  await appContext.close();

  const app = await NestFactory.createMicroservice<MicroserviceOptions>(UserModule, {
    transport: Transport.GRPC,
    options: {
      url: config.url,
      package: 'user',
      protoPath: join(process.cwd(), 'protobufs/user.proto'),
    },
  });

  await app.listen();
}

bootstrap();
