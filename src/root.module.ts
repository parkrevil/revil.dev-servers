import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { TypeOrmModule, TypeOrmModuleOptions } from '@nestjs/typeorm';
import { SnakeNamingStrategy } from 'typeorm-naming-strategies';

import { configs } from './core/configs';
import { UserModule } from './user/user.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
      load: configs,
    }),
    TypeOrmModule.forRootAsync({
      useFactory: (configService: ConfigService) => {
        const config = configService.get<TypeOrmModuleOptions>('typeorm');

        return Object.assign(config, {
          namingStrategy: new SnakeNamingStrategy(),
          autoLoadEntities: true,
          extra: {
            /*             typeCast: (field, next) => {
              const { type } = field;

              if (type === 'DATE' || type === 'DATETIME') {
                const val = field.string();

                return val === null ? null : DateTime.fromSQL(val);
              } else {
                return next();
              }
            },
 */
          },
        });
      },
      inject: [ConfigService],
    }),
    UserModule,
  ],
  controllers: [],
  providers: [],
})
export class RootModule {}
