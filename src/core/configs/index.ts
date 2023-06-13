import typeorm from './typeorm.config';
import server from './server.config';

export const configs = [server, typeorm];
export * from './interfaces';
