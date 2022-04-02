import { CanActivate, ExecutionContext, Injectable, UnauthorizedException, HttpStatus } from '@nestjs/common';
import { RequestModel } from 'src/common/requestModel';
import { AuthService } from './auth.service';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(public readonly authService: AuthService) {}

  public async canActivate(context: ExecutionContext): Promise<boolean> {
    const req: RequestModel = context.switchToHttp().getRequest();
    const authorizationHeader: string = req.headers['authorization'];

    if (!authorizationHeader) {
      throw new UnauthorizedException('Missing Authorization Header');
    }

    const bearer = authorizationHeader.split(' ');

    if (!bearer || bearer.length < 2) {
      throw new UnauthorizedException('Invalid Authorization Header');
    }

    const token = bearer[1];

    const { status, userId } = await this.authService.validate(token);

    req.userId = userId;

    if (status !== HttpStatus.OK) {
      throw new UnauthorizedException('Invalid Token');
    }

    return true;
  }
}
