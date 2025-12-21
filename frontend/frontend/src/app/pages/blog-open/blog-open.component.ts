import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { Observable } from 'rxjs';
@Component({
  selector: 'app-blog-open',
  templateUrl: './blog-open.component.html',
  styleUrls: ['./blog-open.component.css']
})
export class BlogOpenComponent {
  constructor(private AuthService: AuthService) { }
  selectedContent!: Observable<any>;
  title: string = '';
  body: string = '';

  ngOnInit() {
    this.selectedContent = this.AuthService.getCurrentContent();
    this.selectedContent.subscribe((data: any) => {
      if (data) {
        this.title = data.title;
        this.body = data.body;
      }
    });
    console.log('Selected Content Observable subscribed');
  }
}
