import { NestFactory } from '@nestjs/core';
import { RootModule } from './root.module';

export const bootstrap = async () => {
  const app = await NestFactory.create(RootModule);

  await app.listen(3000);
}
