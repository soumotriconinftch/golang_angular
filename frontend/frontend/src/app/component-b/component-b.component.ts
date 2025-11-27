import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { Router } from '@angular/router';
import { UserService, User } from '../services/user.service';

@Component({
  selector: 'app-component-b',
  templateUrl: './component-b.component.html',
  styleUrls: ['./component-b.component.css']
})
export class ComponentBComponent implements OnInit {
  users: User[] = [];
  isLoading = true;
  errorMessage = '';
  username = '';

  constructor(
    private location: Location,
    private router: Router,
    private userService: UserService
  ) {
    const navigation = this.router.getCurrentNavigation();
    this.username = navigation?.extras?.state?.['username'] || '';
  }

  ngOnInit(): void {
    // Check username 'szoumo'
    if (this.username !== 'szoumo') {

      this.router.navigate(['/unauthorized']);
      return;
    }

    this.fetchUsers();
  }

  fetchUsers(): void {
    this.isLoading = true;
    this.errorMessage = '';

    this.userService.getAllUsers(this.username).subscribe({
      next: (users) => {
        this.users = users;
        this.isLoading = false;
      },
      error: (error) => {
        console.error('Error fetching users:', error);
        this.errorMessage = 'Failed to load users. Please try again.';
        this.isLoading = false;
      }
    });
  }

  goBack(): void {
    this.location.back();
  }
}
