import { ArgsType, Field } from '@nestjs/graphql';

@ArgsType()
export class SignInWithEmailArgs {
  @Field()
  email: string;

  @Field()
  password: string;
}
