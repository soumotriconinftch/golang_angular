import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, map, catchError, throwError } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class AuthService {
    private currentUserSubject = new BehaviorSubject<any>(null);
    public currentUser$ = this.currentUserSubject.asObservable();
    private apiUrl = 'http://localhost:8080/user';
    public currentContentSubject = new BehaviorSubject<any>(null);
    public currentContent$ = this.currentContentSubject.asObservable();
    contents:any[] = [];

    constructor(private http: HttpClient) { }

    signup(user: any): Observable<any> {
        return this.http.post(`${this.apiUrl}/sign-up`, user, { withCredentials: true }).pipe(
            map((response) => {
                console.log('User signed up:', response);
                this.showMessage('Signup successful! Please login.');
                return response;
            }),
            catchError((error) => {
                console.error('Signup error:', error);
                this.showMessage('Signup failed: ' + (error.error?.message || error.message || 'Unknown error'));
                return throwError(() => error);
            })
        );
    }

    login(user: any): Observable<any> {
        return this.http.post(`${this.apiUrl}/sign-in`, user, { withCredentials: true }).pipe(
            map((response: any) => {
                console.log('User logged in:', response);
                this.currentUserSubject.next(response);
                localStorage.setItem('user_session', 'true');
                this.showMessage('Login successful!');
                return response;
            }),
            catchError((error) => {
                console.error('Login error:', error);
                this.showMessage('Login failed: ' + (error.error?.message || error.message || 'Unknown error'));
                return throwError(() => error);
            })
        );
    }

    logout(): void {
        localStorage.removeItem('user_session');
        this.currentUserSubject.next(null);
        // We might want to call a backend logout endpoint here if it exists to clear cookies
    }

    isLoggedIn(): boolean {
        return !!localStorage.getItem('user_session');
    }

    private showMessage(message: string): void {
        window.alert(message);
    }

    fetchCurrentUser(): Observable<any> {
        return this.http.get(`${this.apiUrl}/me/content`, { withCredentials: true }).pipe(
            map((user) => {
                console.log("in fetch user", user)
                this.currentUserSubject.next(user);
                return user;
            }),
            catchError((error) => {
                console.error('Failed to fetch user:', error);
                return throwError(() => error);
            })
        );
    }
    fetchUserContent(): Observable<any> {
        return this.http.get(`${this.apiUrl}/me/content`, { withCredentials: true }).pipe(
            map((content) => {
                this.contents = content as any[];
                this.currentContentSubject.next(this.contents);
                return this.contents;
            }),
            catchError((error) => {
                console.error('Failed to fetch user content:', error);
                return throwError(() => error);
            })
        );
    }
    getCurrentContent(): any{
        return this.currentContentSubject.value;
    }
    getCurrentUser(): any {
        return this.currentUserSubject.value;
    }
}
