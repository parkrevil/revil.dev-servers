import apiGateway from './api-gateway.config';
import app from './app.config';
import auth from './auth.config';
import databases from './databases.config';
import user from './user.config';

export const configs = [app, apiGateway, auth, user, databases];
export * from './interfaces';
export * from './types';
