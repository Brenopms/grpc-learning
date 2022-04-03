import { IsEmail, IsString, MaxLength, MinLength } from 'class-validator';
import { LoginRequest, RegisterRequest, ValidateRequest } from './auth.pb';

export class LoginRequestDto implements LoginRequest {
  @IsEmail()
  @MaxLength(50)
  public readonly username: string;

  @IsString()
  @MaxLength(100)
  public readonly password: string;
}

export class RegisterRequestDto implements RegisterRequest {
  @IsEmail()
  public readonly email: string;

  @IsString()
  @MinLength(8)
  @MaxLength(100)
  public readonly password: string;

  @IsString()
  @MaxLength(50)
  public readonly username: string;
}

export class ValidateRequestDto implements ValidateRequest {
  @IsString()
  public readonly token: string;
}
