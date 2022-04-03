import { Injectable } from '@nestjs/common';
import { JwtService as Jwt } from '@nestjs/jwt';
import { scrypt, timingSafeEqual, randomBytes } from 'crypto';
import { promisify } from 'util';
import { User } from '../auth.entity';
import { UserRepository } from '../repository/user.repository';

const scryptAsync = promisify(scrypt);
const SALT_LENGTH = 16;

@Injectable()
export class JwtService {
  constructor(private readonly userRepository: UserRepository, private readonly jwt: Jwt) {}

  public async decode(token: string): Promise<unknown> {
    return this.jwt.decode(token, null);
  }

  public async validateUser(decoded: any) {
    return await this.userRepository.findOne(decoded?.id);
  }

  public generateToken(user: Pick<User, 'id' | 'username' | 'email'>): string {
    return this.jwt.sign({ id: user.id, email: user.email, username: user.username });
  }

  public async isPasswordValid(password: User['password'], requestPassword: User['password']): Promise<boolean> {
    const [salt, key] = password.split(':');
    const keyBuffer = Buffer.from(key, 'hex');

    const derivedKey = (await scryptAsync(requestPassword, salt, 64)) as NodeJS.ArrayBufferView;

    return timingSafeEqual(keyBuffer, derivedKey);
  }

  public async encodePassword(password: User['password']): Promise<string> {
    const salt = randomBytes(SALT_LENGTH).toString('hex');
    const hash = scryptAsync(password, salt, 64);
    return `${salt}:${hash}`;
  }

  public async verify(token: string): Promise<any> {
    return await this.jwt.verify(token);
  }
}
