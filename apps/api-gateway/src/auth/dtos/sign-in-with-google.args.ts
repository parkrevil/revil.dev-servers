import { ArgsType, Field } from '@nestjs/graphql';

@ArgsType()
export class SignInWithGoogleArgs {
  @Field()
  accessToken: string;
}
