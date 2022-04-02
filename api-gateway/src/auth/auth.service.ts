import { Injectable } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { firstValueFrom } from 'rxjs';
import { AuthServiceClient, AUTH_SERVICE_NAME, ValidateResponse } from './auth.pb';

@Injectable()
export class AuthService {
  constructor(private authServiceClient: AuthServiceClient, private readonly client: ClientGrpc) {}

  public onModuleInit(): void {
    this.authServiceClient = this.client.getClientByServiceName<AuthServiceClient>(AUTH_SERVICE_NAME);
  }

  public async validate(token: string): Promise<ValidateResponse> {
    return firstValueFrom(this.authServiceClient.validate({ token }));
  }
}
