import { registerAs } from '@nestjs/config';

import { DatabasesConfig } from './interfaces';

export default registerAs(
  'databases',
  (): DatabasesConfig => ({
    mongodbUri: process.env.MONGODB_URI,
    redis: {
      host: process.env.REDIS_HOST,
      port: parseInt(process.env.REDIS_PORT),
      password: process.env.REDIS_PASSWORD,
    },
  }),
);
