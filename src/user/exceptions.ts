import { BadRequestException } from '@nestjs/common';

export class EmailAlreadyExistsException extends BadRequestException {}
