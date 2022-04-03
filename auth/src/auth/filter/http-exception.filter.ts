import { ArgumentsHost, Catch, ExceptionFilter, HttpException, HttpStatus } from '@nestjs/common';
import { Response } from 'express';

/**
 * Filter to avoid Class-validator throwing HTTP errors into grpc server
 */
@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter {
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const res: Response = ctx.getResponse();
    const req: Request = ctx.getRequest();
    const status: HttpStatus = exception.getStatus();

    if (status === HttpStatus.BAD_REQUEST) {
      const exceptionResponse: any = exception.getResponse();

      return { status, error: exceptionResponse.message };
    }

    res.status(status).json({
      statusCode: status,
      timestamp: new Date().toISOString(),
      path: req.url,
    });
  }
}
