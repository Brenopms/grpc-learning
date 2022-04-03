import { HttpStatus, Injectable } from '@nestjs/common';
import { LoginRequestDto, RegisterRequestDto } from '../auth.dto';
import { User } from '../auth.entity';
import { LoginResponse, RegisterResponse } from '../auth.pb';
import { UserRepository } from '../repository/user.repository';
import { JwtService } from './jwt.service';

@Injectable()
export class AuthService {
  constructor(private readonly userRepository: UserRepository, private readonly jwtService: JwtService) {}

  public async register({ username, email, password }: RegisterRequestDto): Promise<RegisterResponse> {
    const user = await this.userRepository.findByEmail(email);

    if (user) {
      return { status: HttpStatus.CONFLICT, error: ['Email already exists'] };
    }

    const hashedPassword = await this.jwtService.encodePassword(password);
    const newUser: RegisterRequestDto = { username, email, password: hashedPassword };

    await this.userRepository.insertUser(newUser);

    return { status: HttpStatus.CREATED, error: null };
  }

  public async login({ username, password }: LoginRequestDto): Promise<LoginResponse> {
    const user: User = await this.userRepository.findByUsername(username);

    if (!user) {
      return { status: HttpStatus.NOT_FOUND, error: ['Invalid Username or Password'], token: null };
    }

    const isPasswordValid = await this.jwtService.isPasswordValid(user.password, password);
    if (!isPasswordValid) {
      return { status: HttpStatus.NOT_FOUND, error: ['Invalid Username or Password'], token: null };
    }

    const token = this.jwtService.generateToken(user);

    return { token, status: HttpStatus.OK, error: null };
  }
}
