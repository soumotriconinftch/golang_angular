import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-component-d',
  templateUrl: './component-d.component.html',
  styleUrls: ['./component-d.component.css']
})
export class ComponentDComponent {
  constructor(private router: Router) { }

  goBack() {
    this.router.navigate(['/']);
  }
}
