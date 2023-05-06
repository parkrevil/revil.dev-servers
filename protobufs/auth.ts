/* eslint-disable */
import { Metadata } from '@grpc/grpc-js';
import { GrpcMethod, GrpcStreamMethod } from '@nestjs/microservices';
import { Observable } from 'rxjs';
import { BoolValue } from '../google/protobuf/wrappers';

export const protobufPackage = 'auth';

export interface SignInParams {
  username: string;
  password: string;
}

export interface VerifyAccessTokenParams {
  accessToken: string;
}

export interface AuthTokens {
  accessToken: string;
  refreshToken: string;
}

export const AUTH_PACKAGE_NAME = 'auth';

export interface AuthServiceClient {
  signIn(request: SignInParams, metadata?: Metadata): Observable<AuthTokens>;

  verifyAccessToken(request: VerifyAccessTokenParams, metadata?: Metadata): Observable<BoolValue>;
}

export interface AuthServiceController {
  signIn(request: SignInParams, metadata?: Metadata): Promise<AuthTokens> | Observable<AuthTokens> | AuthTokens;

  verifyAccessToken(
    request: VerifyAccessTokenParams,
    metadata?: Metadata,
  ): Promise<BoolValue> | Observable<BoolValue> | BoolValue;
}

export function AuthServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ['signIn', 'verifyAccessToken'];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod('AuthService', method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod('AuthService', method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const AUTH_SERVICE_NAME = 'AuthService';
