import { ApiGatewayConfig } from '@app/core/configs';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';

import { ApiGatewayModule } from './api-gateway.module';

async function bootstrap() {
  const app = await NestFactory.create(ApiGatewayModule);
  const apiGatewayConfig = app.get(ConfigService).get<ApiGatewayConfig>('apiGateway');

  await app.listen(apiGatewayConfig.port, apiGatewayConfig.host);
}
bootstrap();
