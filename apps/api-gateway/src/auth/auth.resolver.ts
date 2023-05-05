import { Args, Query, Resolver } from '@nestjs/graphql';

import { SignInWithEmailArgs, SignInWithGoogleArgs } from './dtos';
import { Auth } from './models';

@Resolver(() => Auth)
export class AuthResolver {
  @Query(() => Auth)
  signWithEmail(@Args() args: SignInWithEmailArgs): Auth {
    const res = new Auth();
    res.accessToken = args.email;
    res.refreshToken = args.password;

    return res;
  }

  @Query(() => Auth)
  signInWIthGoogle(@Args() args: SignInWithGoogleArgs): Auth {
    const res = new Auth();
    res.accessToken = 'a';
    res.refreshToken = 'b';

    return res;
  }
}
