import { instanceToPlain } from 'class-transformer';
import { Observable } from 'rxjs';
import { map } from 'rxjs';
import {
    ExecutionContext,
    NestInterceptor,
    CallHandler,
    Injectable,
} from '@nestjs/common';

@Injectable()
export class TransformInterceptor implements NestInterceptor {
    public intercept(
        context: ExecutionContext,
        next: CallHandler
    ): Observable<any> {
        return next.handle().pipe(map((data) => instanceToPlain(data)));
    }
}
