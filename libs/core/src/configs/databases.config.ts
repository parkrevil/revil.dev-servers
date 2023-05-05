import { registerAs } from '@nestjs/config';

import { DatabasesConfig } from './interfaces';

export default registerAs(
  'databases',
  (): DatabasesConfig => ({
    mongodbUri: process.env.MONGODB_URI,
  }),
);
