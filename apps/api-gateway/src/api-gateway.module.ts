import { CoreModule } from '@app/core';
import { ApiGatewayConfig, DatabasesConfig } from '@app/core/configs';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { CacheModule } from '@nestjs/cache-manager';
import { Module } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { GraphQLModule } from '@nestjs/graphql';
import { MongooseModule } from '@nestjs/mongoose';
import { redisStore } from 'cache-manager-redis-yet';
import { join } from 'path';
import type { RedisClientOptions } from 'redis';

import { AuthModule } from './auth';

@Module({
  imports: [
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      autoSchemaFile: join(process.cwd(), 'apps/api-gateway/src/schema.gql'),
      playground: true,
    }),
    MongooseModule.forRootAsync({
      useFactory: (configService: ConfigService) => ({
        uri: configService.get<DatabasesConfig>('databases').mongodbUri,
      }),
      inject: [ConfigService],
    }),
    CacheModule.registerAsync<RedisClientOptions>({
      useFactory: async (configService: ConfigService) => {
        const apiGatewayConfig = configService.get<ApiGatewayConfig>('apiGateway');
        const databasesConfig = configService.get<DatabasesConfig>('databases');

        return {
          store: await redisStore({
            socket: {
              host: databasesConfig.redis.host,
              port: databasesConfig.redis.port,
            },
            password: databasesConfig.redis.password,
            database: apiGatewayConfig.redisDb,
          }),
          ttl: 365 * 24 * 60 * 60 * 1000,
        };
      },
      inject: [ConfigService],
    }),
    CoreModule,
    AuthModule,
  ],
})
export class ApiGatewayModule {}
