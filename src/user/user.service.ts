import { Injectable } from '@nestjs/common';
import { CreateUserDto } from './dtos';
import { InjectRepository } from '@nestjs/typeorm';
import { User } from './user.entity';
import { Repository } from 'typeorm';

@Injectable()
export class UserService {
  constructor(
    @InjectRepository(User)
    private userRepo: Repository<User>,
  ) {}

  async hasUsername(username: string): Promise<boolean> {
    return (await this.userRepo.countBy({ username })) > 0;
  }

  async create(params: CreateUserDto): Promise<User> {
    const user = new User(params);

    await this.userRepo.insert(user);

    return user;
  }
}
