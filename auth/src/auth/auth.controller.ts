import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { LoginRequestDto, RegisterRequestDto, ValidateRequestDto } from './auth.dto';
import { AUTH_SERVICE_NAME, LoginResponse, RegisterResponse, ValidateResponse } from './auth.pb';
import { AuthService } from './service/auth.service';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @GrpcMethod(AUTH_SERVICE_NAME, 'Register')
  private async register(payload: RegisterRequestDto): Promise<RegisterResponse> {
    return await this.authService.register(payload);
  }

  @GrpcMethod(AUTH_SERVICE_NAME, 'Login')
  private async login(payload: LoginRequestDto): Promise<LoginResponse> {
    return await this.authService.login(payload);
  }

  @GrpcMethod(AUTH_SERVICE_NAME, 'Validate')
  private async validate(payload: ValidateRequestDto): Promise<ValidateResponse> {
    return await this.authService.validate(payload);
  }
}
