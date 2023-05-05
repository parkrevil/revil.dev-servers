import { registerAs } from '@nestjs/config';

import { AuthConfig } from './interfaces';

export default registerAs(
  'auth',
  (): AuthConfig => ({
    host: process.env.AUTH_HOST,
    port: parseInt(process.env.AUTH_PORT),
  }),
);
