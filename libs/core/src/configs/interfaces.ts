import { Env } from '../enums';

export interface AppConfig {
  env: Env;
}

export interface GrpcServerConfig {
  url: string;
}

export interface RedisConfig {
  host: string;
  port: number;
  password: string;
}

export interface RedisDatabaseConfig {
  redisDb: number;
}

export interface ApiGatewayConfig extends RedisDatabaseConfig {
  host: string;
  port: number;
}

export interface DatabasesConfig {
  mongodbUri: string;
  redis: RedisConfig;
}
