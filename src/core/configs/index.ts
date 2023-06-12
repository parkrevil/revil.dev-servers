import mongo from './mongo.config';
import server from './server.config';

export const configs = [server, mongo];
export * from './interfaces';
