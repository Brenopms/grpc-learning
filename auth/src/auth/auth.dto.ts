import { IsEmail, IsString, MaxLength, MinLength } from 'class-validator';
import { User } from './auth.entity';
import { LoginRequest, RegisterRequest, ValidateRequest } from './auth.pb';

export class LoginRequestDto implements LoginRequest {
  @IsEmail()
  @MaxLength(50)
  public readonly username: User['username'];

  @IsString()
  @MaxLength(100)
  public readonly password: User['password'];
}

export class RegisterRequestDto implements RegisterRequest {
  @IsEmail()
  public readonly email: User['email'];

  @IsString()
  @MinLength(8)
  @MaxLength(100)
  public readonly password: User['password'];

  @IsString()
  @MaxLength(50)
  public readonly username: User['username'];
}

export class ValidateRequestDto implements ValidateRequest {
  @IsString()
  public readonly token: string;
}
