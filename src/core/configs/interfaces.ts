import { CorsOptions } from '@nestjs/common/interfaces/external/cors-options.interface';
import { Env } from '../enums';

export interface ServerConfig {
  env: Env;
  listen: {
    host: string;
    port: number;
  };
  cors: CorsOptions;
}

export interface MongoConfig {
  uri: string;
  db: string;
}
