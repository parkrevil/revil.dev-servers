import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { configs } from './configs';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      load: configs,
    }),
  ],
  exports: [],
})
export class CoreModule {}
