import { CoreModule } from '@app/core';
import { AuthConfig, DatabasesConfig } from '@app/core/configs';
import { CacheModule } from '@nestjs/cache-manager';
import { Module } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { redisStore } from 'cache-manager-redis-yet';
import type { RedisClientOptions } from 'redis';

import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';

@Module({
  imports: [
    CacheModule.registerAsync<RedisClientOptions>({
      useFactory: async (configService: ConfigService) => {
        const authConfig = configService.get<AuthConfig>('auth');
        const databasesConfig = configService.get<DatabasesConfig>('databases');

        return {
          store: await redisStore({
            socket: {
              host: databasesConfig.redis.host,
              port: databasesConfig.redis.port,
            },
            password: databasesConfig.redis.password,
            database: authConfig.redisDb,
          }),
          ttl: 365 * 24 * 60 * 60 * 1000,
        };
      },
      inject: [ConfigService],
    }),
    CoreModule,
  ],
  controllers: [AuthController],
  providers: [AuthService],
})
export class AuthModule {}
