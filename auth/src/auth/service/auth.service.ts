import { HttpStatus, Injectable } from '@nestjs/common';
import { RegisterRequestDto } from '../auth.dto';
import { RegisterResponse } from '../auth.pb';
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
}
