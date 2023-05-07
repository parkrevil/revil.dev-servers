import { registerAs } from '@nestjs/config';

import { AuthConfig } from './types';

export default registerAs(
  'auth',
  (): AuthConfig => ({
    url: process.env.AUTH_URL,
    redisDb: parseInt(process.env.AUTH_REDIS_DB),
  }),
);
