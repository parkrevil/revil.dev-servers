import { Env } from '../enums';

export interface AppConfig {
  env: Env;
}

interface BaseServerConfig {
  host: string;
  port: number;
}

export type ApiGatewayConfig = BaseServerConfig;

export type UserConfig = BaseServerConfig;

export type AuthConfig = BaseServerConfig;
