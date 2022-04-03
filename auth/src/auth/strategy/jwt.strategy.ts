import { Inject } from '@nestjs/common';
import { PassportStrategy } from '@nestjs/passport';
import { ExtractJwt, Strategy } from 'passport-jwt';
import { User } from '../auth.entity';
import { JwtService } from '../service/jwt.service';

export class JwtStrategy extends PassportStrategy(Strategy) {
  @Inject(JwtService)
  private readonly jwtService: JwtService;

  constructor() {
    super({ jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(), secretOrKey: 'dev', ignoreExpiration: true });
  }

  private validate(token: string): Promise<User | never> {
    return this.jwtService.validateUser(token);
  }
}
