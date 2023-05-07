import { GrpcServerConfig, RedisDatabaseConfig } from './interfaces';

export type UserConfig = GrpcServerConfig & RedisDatabaseConfig;

export type AuthConfig = GrpcServerConfig & RedisDatabaseConfig;
