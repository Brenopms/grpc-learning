import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { ORDER_PACKAGE_NAME, ORDER_SERVICE_NAME } from 'src/order/order.pb';
import { ProductController } from './product.controller';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: ORDER_SERVICE_NAME,
        transport: Transport.GRPC,
        options: {
          url: '0.0.0.0:50053',
          package: ORDER_PACKAGE_NAME,
          protoPath: '../grpc-proto/proto/product.proto',
        },
      },
    ]),
  ],
  controllers: [ProductController],
})
export class ProductModule {}
