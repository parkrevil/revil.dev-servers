import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';

import { CreateWithUsernameDto } from './dtos';
import { User } from './schemas';

@Injectable()
export class UserService {
  constructor(@InjectModel(User.name) private userModel: Model<User>) {}

  createWithUsername(params: CreateWithUsernameDto): Promise<User> {
    const user = new this.userModel(params);

    return user.save();
  }
}
