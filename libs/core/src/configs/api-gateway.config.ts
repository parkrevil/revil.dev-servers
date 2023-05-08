import { registerAs } from '@nestjs/config';

import { ApiGatewayConfig } from './interfaces';

export default registerAs(
  'apiGateway',
  (): ApiGatewayConfig => ({
    host: process.env.API_GATEWAY_HOST,
    port: parseInt(process.env.API_GATEWAY_PORT),
    cors: {
      origin: process.env.API_GATEWAY_CORS_ORIGIN.split(',').map((regexp) => {
        return new RegExp(regexp);
      }),
      methods: process.env.API_GATEWAY_CORS_METHODS.split(','),
      allowedHeaders: process.env.API_GATEWAY_CORS_ALLOWED_HEADERS.split(','),
      exposedHeaders: process.env.API_GATEWAY_CORS_EXPOSED_HEADERS.split(','),
      preflightContinue: process.env.API_GATEWAY_CORS_PREFLIGHT_CONTINUE === 'true',
      credentials: process.env.API_GATEWAY_CORS_CREDENTIALS === 'true',
      maxAge: 3600,
    },
    redisDb: parseInt(process.env.API_GATEWAY_REDIS_DB),
  }),
);
