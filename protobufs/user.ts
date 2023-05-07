/* eslint-disable */
import { Metadata } from '@grpc/grpc-js';
import { GrpcMethod, GrpcStreamMethod } from '@nestjs/microservices';
import { Observable } from 'rxjs';

export const protobufPackage = 'user';

export interface UserId {
  id: string;
}

export interface createUserWithUsernameParams {
  username: string;
  password: string;
}

export const USER_PACKAGE_NAME = 'user';

export interface UserServiceClient {
  createUserWithUsername(request: createUserWithUsernameParams, metadata?: Metadata): Observable<UserId>;
}

export interface UserServiceController {
  createUserWithUsername(
    request: createUserWithUsernameParams,
    metadata?: Metadata,
  ): Promise<UserId> | Observable<UserId> | UserId;
}

export function UserServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ['createUserWithUsername'];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod('UserService', method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod('UserService', method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const USER_SERVICE_NAME = 'UserService';
