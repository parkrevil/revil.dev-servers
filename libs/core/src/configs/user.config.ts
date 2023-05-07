import { registerAs } from '@nestjs/config';

import { UserConfig } from './types';

export default registerAs(
  'user',
  (): UserConfig => ({
    url: process.env.USER_URL,
    redisDb: parseInt(process.env.USER_REDIS_DB),
  }),
);
