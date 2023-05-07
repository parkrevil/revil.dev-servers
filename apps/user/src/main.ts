import { UserConfig } from '@app/core/configs';
import { ClassSerializerInterceptor, ValidationPipe } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { NestFactory, Reflector } from '@nestjs/core';
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

  app
    .useGlobalInterceptors(
      new ClassSerializerInterceptor(app.get(Reflector), {
        strategy: 'exposeAll',
        excludeExtraneousValues: true,
      }),
    )
    .useGlobalPipes(new ValidationPipe());
  await app.listen();
}

bootstrap();
