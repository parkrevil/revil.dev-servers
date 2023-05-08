import { isLocal } from '@app/core';
import { ApiGatewayConfig } from '@app/core/configs';
import { ConfigService } from '@nestjs/config';
import { NestFactory } from '@nestjs/core';
import * as compression from 'compression';
import * as cookieParser from 'cookie-parser';
import helmet, { HelmetOptions } from 'helmet';

import { ApiGatewayModule } from './api-gateway.module';

async function bootstrap() {
  const app = await NestFactory.create(ApiGatewayModule);
  const apiGatewayConfig = app.get(ConfigService).get<ApiGatewayConfig>('apiGateway');

  app.enableCors(apiGatewayConfig.cors);
  app.use(compression());
  app.use(cookieParser());

  const helmetOptions: HelmetOptions = {};

  if (isLocal) {
    helmetOptions.crossOriginEmbedderPolicy = false;
    helmetOptions.contentSecurityPolicy = {
      directives: {
        imgSrc: [`'self'`, 'data:', 'apollo-server-landing-page.cdn.apollographql.com'],
        scriptSrc: [`'self'`, `https: 'unsafe-inline'`],
        manifestSrc: [`'self'`, 'apollo-server-landing-page.cdn.apollographql.com'],
        frameSrc: [`'self'`, 'sandbox.embed.apollographql.com'],
      },
    };
  }

  app.use(helmet(helmetOptions));

  await app.listen(apiGatewayConfig.port, apiGatewayConfig.host);
}
bootstrap();
