import { Controller, Post, Body, Inject } from '@nestjs/common';
import { CreateUsersDto } from './create-users.dto';
import { UsersService } from './users.service';

@Controller('users')
export class UsersController {
  constructor(@Inject(UsersService) private usersService: UsersService) {}

  @Post()
  async create(@Body() body: CreateUsersDto) {
    return this.usersService.create(body);
  }

  // @Get()
  // getAll(): Observable<[UsersModel]> {
  //   return this.usersService.getAll();
  // }
}
