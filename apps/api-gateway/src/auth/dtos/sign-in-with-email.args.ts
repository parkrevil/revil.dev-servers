import { Field, ArgsType } from '@nestjs/graphql';

@ArgsType()
export class SignInWithEmailArgs {
  @Field()
  email: string;

  @Field()
  password: string;
}
