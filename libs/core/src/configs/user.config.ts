import { registerAs } from '@nestjs/config';

import { UserConfig } from './interfaces';

export default registerAs(
  'user',
  (): UserConfig => ({
    host: process.env.USER_HOST,
    port: parseInt(process.env.USER_PORT),
  }),
);
