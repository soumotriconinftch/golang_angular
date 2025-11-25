import { Component } from '@angular/core';
import { Location } from '@angular/common';

@Component({
  selector: 'app-component-b',
  templateUrl: './component-b.component.html',
  styleUrls: ['./component-b.component.css']
})
export class ComponentBComponent {

  constructor(private location: Location) {}

  goBack(): void {
    this.location.back();
  }

}
