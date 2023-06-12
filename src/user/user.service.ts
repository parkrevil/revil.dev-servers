import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { User, UserDocument } from './schemas';
import { Model } from 'mongoose';
import { CreateUserDto } from './dtos';

@Injectable()
export class UserService {
  constructor(@InjectModel(User.name) private userModel: Model<UserDocument>) {}

  async hasEmail(email: string): Promise<boolean> {
    return !!(await this.userModel.exists({ email }));
  }

  create(params: CreateUserDto): Promise<User> {
    const user = new this.userModel(params);

    return user.save();
  }
}
