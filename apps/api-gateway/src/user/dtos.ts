import { Field, InputType } from '@nestjs/graphql';

@InputType()
export class CreateUserWithUsernameInput {
  @Field()
  username: string;

  @Field()
  password: string;

  @Field()
  nickname: string;
}
