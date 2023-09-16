import { HttpStatus } from '@nestjs/common';
import { GrpcStatus } from './grpc-status.enum';

export function convertGrpcErrorTo(code: GrpcStatus): HttpStatus {
  switch (code) {
    case GrpcStatus.Canceled:
      return HttpStatus.NOT_ACCEPTABLE;
    default:
      return HttpStatus.BAD_REQUEST;
  }
}
