import { NestFactory } from '@nestjs/core';
import {
  FastifyAdapter,
  NestFastifyApplication,
} from '@nestjs/platform-fastify';
import compression from '@fastify/compress';
import { ConfigService } from '@nestjs/config';
import { RootModule } from './root.module';
import { ServerConfig } from './core/configs';
import { isLocal } from './core/helpers';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';

export const bootstrap = async () => {
  const app = await NestFactory.create<NestFastifyApplication>(RootModule, new FastifyAdapter());
  const serverConfig = app.get(ConfigService).get<ServerConfig>('server');
  
  app.enableCors(serverConfig.cors);
  await app.register(compression);

  if (isLocal()) {
    const apiDoc = new DocumentBuilder()
      .setTitle('revil.dev')
      .setVersion('0.0')
      .addBearerAuth()
      .addTag('인증')
      .addTag('사용자')
      .build();
    const document = SwaggerModule.createDocument(app, apiDoc);

    SwaggerModule.setup('docs', app, document, {
      swaggerOptions: {
        docExpansion: 'none',
      },
    });
  }

  await app.listen(serverConfig.listen.port, serverConfig.listen.host);
}