import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { PrismaService } from './prisma.service';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  /* 
    Prisma interferes with NestJS enableShutdownHooks.
    Prisma listens for shutdown signals and will call process.exit() before your application shutdown hooks fire.
    To deal with this, you would need to add a listener for Prisma beforeExit event.
   */
  const prismaService = app.get(PrismaService);
  await prismaService.enableShutdownHooks(app);

  await app.listen(3000);
}
bootstrap();
