import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, tap } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class AuthService {
    private currentUserSubject = new BehaviorSubject<any>(null);
    public currentUser$ = this.currentUserSubject.asObservable();
    private apiUrl = 'http://localhost:9000/user';

    constructor(private http: HttpClient) { }

    signup(user: any): Observable<any> {
        return this.http.post(`${this.apiUrl}/sign-up`, user).pipe(
            tap({
                next: (response) => {
                    console.log('User signed up:', response);
                    this.showMessage('Signup successful! Please login.');
                },
                error: (error) => {
                    console.error('Signup error:', error);
                    this.showMessage('Signup failed: ' + (error.error?.message || error.message || 'Unknown error'));
                }
            })
        );
    }

    login(user: any): Observable<any> {
        return this.http.post(`${this.apiUrl}/sign-in`, user).pipe(
            tap({
                next: (response: any) => {
                    console.log('User logged in:', response);
                    this.currentUserSubject.next(response);
                    this.showMessage('Login successful!');
                },
                error: (error) => {
                    console.error('Login error:', error);
                    this.showMessage('Login failed: ' + (error.error?.message || error.message || 'Unknown error'));
                }
            })
        );
    }

    private showMessage(message: string): void {
        window.alert(message);
    }

    getCurrentUser(): any {
        return this.currentUserSubject.value;
    }
}
