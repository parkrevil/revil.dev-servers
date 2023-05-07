import { grpcClientInjectionTokens } from '@app/core';
import { Inject, OnModuleInit } from '@nestjs/common';
import { Args, Query, Resolver } from '@nestjs/graphql';
import { ClientGrpc } from '@nestjs/microservices';
import { AuthServiceClient } from 'protobufs/auth';
import { Observable } from 'rxjs';

import { SignInWithEmailArgs, SignInWithGoogleArgs } from './dtos';
import { AuthTokens } from './models';

@Resolver()
export class AuthResolver implements OnModuleInit {
  private authServiceClient: AuthServiceClient;

  constructor(@Inject(grpcClientInjectionTokens.auth) private authGrpcClient: ClientGrpc) {}

  onModuleInit(): void {
    this.authServiceClient = this.authGrpcClient.getService<AuthServiceClient>('AuthService');
  }

  @Query(() => AuthTokens)
  signInWithEmail(@Args() args: SignInWithEmailArgs): Observable<AuthTokens> {
    return this.authServiceClient.signIn({
      username: 'username',
      password: 'password',
    });
  }

  @Query(() => AuthTokens)
  signInWithGoogle(@Args() args: SignInWithGoogleArgs): Observable<AuthTokens> {
    return this.authServiceClient.signIn({
      username: 'username',
      password: 'password',
    });
  }
}
