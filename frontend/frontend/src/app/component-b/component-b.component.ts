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
    // Get username from navigation state
    const navigation = this.router.getCurrentNavigation();
    this.username = navigation?.extras?.state?.['username'] || '';
  }

  ngOnInit(): void {
    // Check if username is 'szoumo'
    if (this.username !== 'szoumo') {
      // Not authorized, redirect to ComponentC
      this.router.navigate(['/unauthorized']);
      return;
    }

    // Authorized, fetch users
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
