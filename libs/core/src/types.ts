import { GrpcServerConfig, RedisDatabaseConfig } from './configs';

export type UserConfig = GrpcServerConfig & RedisDatabaseConfig;

export type AuthConfig = GrpcServerConfig & RedisDatabaseConfig;
