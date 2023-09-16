import { Transport, ClientOptions } from '@nestjs/microservices';

export const grpcClientOptions: ClientOptions = {
  transport: Transport.GRPC,
  options: {
    url: 'localhost:5000',
    package: 'core.grpc.api',
    protoPath: ['./proto/users_service.proto'],
  },
};
