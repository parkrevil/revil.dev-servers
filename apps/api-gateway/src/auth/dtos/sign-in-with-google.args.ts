import { Field, ArgsType } from '@nestjs/graphql';

@ArgsType()
export class SignInWithGoogleArgs {
  @Field()
  accessToken: string;
}
