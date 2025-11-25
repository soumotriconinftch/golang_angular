import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-component-c',
  templateUrl: './component-c.component.html',
  styleUrls: ['./component-c.component.css']
})
export class ComponentCComponent {
  constructor(private router: Router) { }

  goBack(): void {
    this.router.navigate(['/']);
  }
}
