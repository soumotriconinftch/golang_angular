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
    // Check "szoumo"
    if (this.user.username.toLowerCase() === 'szoumo') {
      this.router.navigate(['/user'], {
        state: { username: this.user.username }
      });
    } else {
      // everyone else
      this.router.navigate(['/unauthorized']);
    }
  }
}
