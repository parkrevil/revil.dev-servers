import { NestFactory } from '@nestjs/core';
import {
  FastifyAdapter,
  NestFastifyApplication,
} from '@nestjs/platform-fastify';
import { ConfigService } from '@nestjs/config';
import { RootModule } from './root.module';
import { ServerConfig } from './core/configs';

export const bootstrap = async () => {
  const app = await NestFactory.create<NestFastifyApplication>(RootModule, new FastifyAdapter());
  const serverConfig = app.get(ConfigService).get<ServerConfig>('server');
  
  app.enableCors(serverConfig.cors);

  await app.listen(serverConfig.listen.port, serverConfig.listen.host);
}
