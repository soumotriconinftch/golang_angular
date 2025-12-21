import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  username: string = '';

  constructor(private authService: AuthService) { }

  ngOnInit(): void {
    console.log("dashboard component")
    this.authService.currentUser$.subscribe(user => {
      console.log("user", user)
      if (user) {
        this.username = user.Username || user.username; // Handle case sensitivity
      } else if (this.authService.isLoggedIn()) {
        // If logged in but no user data (e.g. refresh), fetch it
        this.authService.fetchCurrentUser().subscribe(x => {
          
          console.log("user", x)
        });
      }
    });
  }
}
