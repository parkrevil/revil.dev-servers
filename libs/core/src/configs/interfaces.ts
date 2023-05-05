import { Env } from '../enums';

export interface AppConfig {
  env: Env;
}

interface BaseServerConfig {
  host: string;
  port: number;
}

export interface ApiGatewayConfig extends BaseServerConfig {}

export interface UserConfig extends BaseServerConfig {}

export interface AuthConfig extends BaseServerConfig {}
