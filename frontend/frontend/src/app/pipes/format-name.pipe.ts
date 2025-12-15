import { Pipe, PipeTransform } from '@angular/core';
import { Observable, timer, map } from 'rxjs';

@Pipe({
    name: 'formatName'
})
export class FormatNamePipe implements PipeTransform {
    transform(value: string): string {
        return value.toUpperCase();
    }
}
