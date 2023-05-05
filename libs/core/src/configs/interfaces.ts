import { Env } from '../enums';

export interface AppConfig {
  env: Env;
}

interface BaseServerConfig {
  host: string;
  port: number;
}

interface RedisConfig {
  host: string;
  port: number;
  password: string;
}

interface RedisDatabaseConfig {
  redisDb: number;
}

export type ApiGatewayConfig = BaseServerConfig & RedisDatabaseConfig;

export type UserConfig = BaseServerConfig & RedisDatabaseConfig;

export type AuthConfig = BaseServerConfig & RedisDatabaseConfig;

export interface DatabasesConfig {
  mongodbUri: string;
  redis: RedisConfig;
}
