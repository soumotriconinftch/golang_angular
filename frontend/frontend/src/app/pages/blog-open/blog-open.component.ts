import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
@Component({
  selector: 'app-blog-open',
  templateUrl: './blog-open.component.html',
  styleUrls: ['./blog-open.component.css']
})
export class BlogOpenComponent {
  constructor(private AuthService: AuthService) {}
  selectedContent: any = '';

  ngOnInit() {
    this.selectedContent = this.AuthService.getCurrentContent();
    console.log('Selected Content:', this.selectedContent);
  }
}
