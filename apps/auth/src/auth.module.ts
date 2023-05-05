import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { CoreModule } from '@app/core';

@Module({
  imports: [
    CoreModule,
  ],
  controllers: [AuthController],
  providers: [AuthService],
})
export class AuthModule {}
