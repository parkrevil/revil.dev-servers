import { CoreModule } from '@app/core';
import { DatabasesConfig } from '@app/core/configs';
import { Module } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';

import { UserController } from './user.controller';
import { UserService } from './user.service';

@Module({
  imports: [
    MongooseModule.forRootAsync({
      useFactory: (configService: ConfigService) => ({
        uri: configService.get<DatabasesConfig>('databases').mongodbUri,
      }),
      inject: [ConfigService],
    }),
    CoreModule,
  ],
  controllers: [UserController],
  providers: [UserService],
})
export class UserModule {}
