import { Env } from './enums';

export const isLocal = () => process.env.NODE_ENV === Env.Local;
export const isProduction = () => process.env.NODE_ENV === Env.Production;
