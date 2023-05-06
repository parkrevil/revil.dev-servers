import { ApiGatewayConfig } from '@app/core/configs';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import * as compression from 'compression';
import * as cookieParser from 'cookie-parser';

import { ApiGatewayModule } from './api-gateway.module';

async function bootstrap() {
  const app = await NestFactory.create(ApiGatewayModule);
  const apiGatewayConfig = app.get(ConfigService).get<ApiGatewayConfig>('apiGateway');

  app.use(compression());
  app.use(cookieParser());

  await app.listen(apiGatewayConfig.port, apiGatewayConfig.host);
}
bootstrap();
