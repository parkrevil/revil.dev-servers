import { NestFactory } from '@nestjs/core';
import { UserConfig } from '@app/core/configs';
import { ConfigService } from '@nestjs/config';
import { UserModule } from './user.module';

async function bootstrap() {
  const app = await NestFactory.create(UserModule);
  const userConfig = app.get(ConfigService).get<UserConfig>('user');

  await app.listen(userConfig.port, userConfig.host);
}
bootstrap();
