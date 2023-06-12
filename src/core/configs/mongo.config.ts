import { registerAs } from '@nestjs/config';

import { MongoConfig } from './interfaces';

export default registerAs(
  'mongo',
  (): MongoConfig => ({
    uri: process.env.MONGO_URI,
    db: process.env.MONGO_DB,
  }),
);
