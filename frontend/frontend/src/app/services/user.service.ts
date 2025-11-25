import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface User {
    id?: number;
    username: string;
    content: string;
}

@Injectable({
    providedIn: 'root'
})
export class UserService {
    private apiUrl = 'http://localhost:8080/users';

    constructor(private http: HttpClient) { }

    createUser(user: User): Observable<User> {
        return this.http.post<User>(this.apiUrl, user);
    }

    getAllUsers(username: string): Observable<User[]> {
        const headers = new HttpHeaders({
            'X-Username': username
        });
        return this.http.get<User[]>(this.apiUrl, { headers });
    }
}
