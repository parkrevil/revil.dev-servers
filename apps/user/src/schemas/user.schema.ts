import { File } from '@app/core/mongoose/schemas';
import { Prop, raw, Schema, SchemaFactory } from '@nestjs/mongoose';
import { DateTime } from 'luxon';
import { HydratedDocument } from 'mongoose';

export type UserDocument = HydratedDocument<User>;

@Schema()
export class User {
  @Prop({
    required: true,
  })
  username: string;

  @Prop({
    required: true,
  })
  password: number;

  @Prop({
    required: true,
  })
  nickname: string;

  @Prop({
    type: File,
    required: false,
  })
  imageUrl: File;

  @Prop({
    type: Date,
    required: true,
  })
  createdAt: DateTime;

  @Prop({
    type: Date,
    required: true,
  })
  updatedAt: DateTime;
}

export const UserSchema = SchemaFactory.createForClass(User);
