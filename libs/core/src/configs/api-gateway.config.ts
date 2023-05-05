import { registerAs } from '@nestjs/config';

import { ApiGatewayConfig } from './interfaces';

export default registerAs(
  'apiGateway',
  (): ApiGatewayConfig => ({
    host: process.env.API_GATEWAY_HOST,
    port: parseInt(process.env.API_GATEWAY_PORT),
    redisDb: parseInt(process.env.API_GATEWAY_REDIS_DB),
  }),
);
