import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-component-a',
  templateUrl: './component-a.component.html',
  styleUrls: ['./component-a.component.css']
})
export class ComponentAComponent {
  constructor(private router: Router) {}

  goToUser() {
    this.router.navigate(['/user']);
  }
}
