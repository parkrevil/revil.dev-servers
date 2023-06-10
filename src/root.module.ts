import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { configs } from './core/configs';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      load: configs,
    }),
  ],
  controllers: [],
  providers: [],
})
export class RootModule {}
