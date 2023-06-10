import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { configs } from './core/configs';
import { UserModule } from './user/user.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      load: configs,
    }),
    UserModule,
  ],
  controllers: [],
  providers: [],
})
export class RootModule {}
