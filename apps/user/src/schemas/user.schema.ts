import { File } from '@app/core/mongoose/schemas';
import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { DateTime } from 'luxon';
import { HydratedDocument } from 'mongoose';

export type UserDocument = HydratedDocument<User>;

@Schema()
export class User {
  @Prop({
    unique: true,
    required: true,
  })
  username: string;

  @Prop({
    required: true,
  })
  password: string;

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
    default: DateTime.local(),
  })
  createdAt: DateTime;

  @Prop({
    type: Date,
    required: true,
    default: DateTime.local(),
  })
  updatedAt: DateTime;
}

export const UserSchema = SchemaFactory.createForClass(User);
