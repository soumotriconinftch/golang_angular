import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { UserService, User } from '../services/user.service';

@Component({
  selector: 'app-component-a',
  templateUrl: './component-a.component.html',
  styleUrls: ['./component-a.component.css']
})
export class ComponentAComponent {
  user: User = {
    username: '',
    content: ''
  };

  isSubmitting = false;
  errorMessage = '';

  constructor(
    private router: Router,
    private userService: UserService
  ) { }

  onSubmit(form: any) {
    if (form.valid) {
      this.isSubmitting = true;
      this.errorMessage = '';

      this.userService.createUser(this.user).subscribe({
        next: (response) => {
          console.log('User created:', response);
          // Navigate to ComponentD (Thank You page with video)
          this.router.navigate(['/thankyou']);
        },
        error: (error) => {
          console.error('Error creating user:', error);
          this.errorMessage = 'Failed to create user. Please try again.';
          this.isSubmitting = false;
        },
        complete: () => {
          this.isSubmitting = false;
        }
      });
    }
  }

  viewUser() {
    // Check if username is "szoumo"
    if (this.user.username.toLowerCase() === 'szoumo') {
      // Navigate to ComponentB for szoumo
      this.router.navigate(['/user'], {
        state: { username: this.user.username }
      });
    } else {
      // Navigate to ComponentC for everyone else
      this.router.navigate(['/unauthorized']);
    }
  }
}
