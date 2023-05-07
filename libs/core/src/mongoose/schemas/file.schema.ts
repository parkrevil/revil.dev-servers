import { Prop, Schema } from '@nestjs/mongoose';

@Schema()
export class File {
  @Prop({
    required: true,
  })
  path: string;

  @Prop({
    required: true,
  })
  name: string;
}
