import { registerAs } from '@nestjs/config';

import { Env } from '../enums';
import { AppConfig } from './interfaces';

export default registerAs(
  'app',
  (): AppConfig => ({
    env: process.env.NODE_ENV as Env,
  }),
);
