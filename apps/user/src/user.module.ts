import { CoreModule } from '@app/core';
import { DatabasesConfig, UserConfig } from '@app/core/configs';
import { CacheModule } from '@nestjs/cache-manager';
import { Module } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';
import { redisStore } from 'cache-manager-redis-yet';
import type { RedisClientOptions } from 'redis';

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
    CacheModule.registerAsync<RedisClientOptions>({
      useFactory: async (configService: ConfigService) => {
        const userConfig = configService.get<UserConfig>('user');
        const databasesConfig = configService.get<DatabasesConfig>('databases');

        return {
          store: await redisStore({
            socket: {
              host: databasesConfig.redis.host,
              port: databasesConfig.redis.port,
            },
            password: databasesConfig.redis.password,
            database: userConfig.redisDb,
          }),
          ttl: 365 * 24 * 60 * 60 * 1000,
        };
      },
      inject: [ConfigService],
    }),
    CoreModule,
  ],
  controllers: [UserController],
  providers: [UserService],
})
export class UserModule {}
