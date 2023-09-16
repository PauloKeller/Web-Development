import { HttpException, HttpStatus } from '@nestjs/common';
import { ClientGrpc, Client } from '@nestjs/microservices';
import { catchError, first, Observable } from 'rxjs';
import { grpcClientOptions } from 'src/grpc-client.options';
import { convertGrpcErrorTo } from 'src/utils/convert.util';
import { CreateUsersDto } from './create-users.dto';
import { UsersModel } from './users.model';

interface UsersServiceClientInterface {
  create(param: CreateUsersDto): Observable<boolean>;
  getAll(): Observable<[UsersModel]>;
  getById(id: string): Observable<UsersModel>;
}

export class UsersService {
  @Client(grpcClientOptions) private readonly client: ClientGrpc;
  private usersService: UsersServiceClientInterface;

  onModuleInit() {
    this.usersService =
      this.client.getService<UsersServiceClientInterface>('UsersService');
  }

  create(param: CreateUsersDto): Observable<boolean | HttpException> {
    return this.usersService.create(param).pipe(
      first(),
      catchError((err) => {
        const httpStatus = convertGrpcErrorTo(err.code);
        const parsedError = JSON.parse(err.details);
        const error = { message: parsedError.message, code: parsedError.code };
        throw new HttpException(error, httpStatus);
      }),
    );
  }
}
