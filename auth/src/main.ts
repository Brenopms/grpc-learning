import { ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { Transport } from '@nestjs/microservices';
import { join } from 'path';
import { AppModule } from './app.module';
import { protobufPackage } from './auth/auth.pb';
import { HttpExceptionFilter } from './auth/filter/http-exception.filter';

async function bootstrap() {
  const app = await NestFactory.createMicroservice(AppModule, {
    transport: Transport.GRPC,
    options: {
      url: '0.0.0.0:50051',
      package: protobufPackage,
      protoPath: join('../grpc-proto/proto/auth.proto'),
    },
  });

  app.useGlobalFilters(new HttpExceptionFilter());
  app.useGlobalPipes(new ValidationPipe({ whitelist: true, transform: true }));

  /* 
    Prisma interferes with NestJS enableShutdownHooks.
    Prisma listens for shutdown signals and will call process.exit() before your application shutdown hooks fire.
    To deal with this, you would need to add a listener for Prisma beforeExit event.
   */
  //const prismaService = app.get(PrismaService);
  //await prismaService.enableShutdownHooks(app);

  await app.listen();
}

bootstrap();
