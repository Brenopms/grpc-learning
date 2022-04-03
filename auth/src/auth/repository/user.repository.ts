import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma.service';
import { RegisterRequestDto } from '../auth.dto';
import { User } from '../auth.entity';

interface IUserRepository {
  findOne(id: User['id']): Promise<User>;
  insertUser(userDto: RegisterRequestDto): Promise<User>;
}

@Injectable()
export class UserRepository implements IUserRepository {
  constructor(private readonly prisma: PrismaService) {}

  async findOne(id: string): Promise<User> {
    const user = await this.prisma.user.findUnique({
      where: {
        id: id,
      },
    });

    return user;
  }

  async findByEmail(email: User['email']): Promise<User> {
    const user = await this.prisma.user.findUnique({
      where: {
        email: email,
      },
    });

    return user;
  }

  async insertUser(userDto: RegisterRequestDto): Promise<User> {
    const newUser = await this.prisma.user.create({
      data: userDto,
    });

    return newUser;
  }
}
