import app from './app.config';
import apiGateway from './api-gateway.config';
import auth from './auth.config';
import user from './user.config';

export const configs = [
  app,
  apiGateway,
  auth,
  user,
];
export * from './interfaces';
