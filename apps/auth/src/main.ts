import { AuthConfig } from '@app/core/configs';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';

import { AuthModule } from './auth.module';

async function bootstrap() {
  const app = await NestFactory.create(AuthModule);
  const authConfig = app.get(ConfigService).get<AuthConfig>('auth');

  await app.listen(authConfig.port, authConfig.host);
}
bootstrap();
