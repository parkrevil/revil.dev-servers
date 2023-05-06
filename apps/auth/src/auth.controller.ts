import { Metadata } from '@grpc/grpc-js';
import { Controller } from '@nestjs/common';
import { BoolValue } from 'google/protobuf/wrappers';
import {
  AuthServiceController,
  AuthServiceControllerMethods,
  AuthTokens,
  SignInParams,
  VerifyAccessTokenParams,
} from 'protobufs/auth';

import { AuthService } from './auth.service';

@Controller()
@AuthServiceControllerMethods()
export class AuthController implements AuthServiceController {
  constructor(private readonly authService: AuthService) {}

  signIn(request: SignInParams, metadata?: Metadata): AuthTokens {
    return {
      accessToken: 'a',
      refreshToken: 'b',
    };
  }

  verifyAccessToken(request: VerifyAccessTokenParams, metadata?: Metadata): BoolValue {
    return {
      value: true,
    };
  }
}
