import { registerAs } from '@nestjs/config';

import { AppConfig } from './interfaces';
import { Env } from '../enums';

export default registerAs(
  'app',
  (): AppConfig => ({
    env: process.env.NODE_ENV as Env,
  }),
);
