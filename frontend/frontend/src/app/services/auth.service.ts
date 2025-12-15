import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class AuthService {
    private currentUserSubject = new BehaviorSubject<any>(null);
    public currentUser$ = this.currentUserSubject.asObservable();

    constructor() { }

    signup(user: any): boolean {
        // Simulate signup
        console.log('User signed up:', user);
        this.currentUserSubject.next(user);
        return true;
    }

    getCurrentUser(): any {
        return this.currentUserSubject.value;
    }
}
