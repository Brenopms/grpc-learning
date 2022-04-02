import { Request } from 'express';
import { ValidateResponse } from 'src/auth/auth.pb';

export interface RequestModel extends Request {
  user?: ValidateResponse['userId'];
}
