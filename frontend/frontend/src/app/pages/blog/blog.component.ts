import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-blog',
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.css']
})
export class BlogComponent {
  constructor(private AuthService: AuthService, private router: Router) { }
  data: any;
  content: any = [];
  ngOnInit() {
    this.AuthService.fetchUserContent().subscribe({
      next: (data: any) => {
        console.log(data);
        this.data = data;
        console.log('User content fetched:', data);
        this.content = data;
      },
      error: () => {
        console.error('Error fetching user content:');
      },
      complete: () => {
        console.log('Fetch user content completed.');
      }
    });
  }
  clicked(item: any) {
    console.log('Button clicked', item);
    this.AuthService.setCurrentContent(item);
    this.router.navigate(['/blog-open']);
  }
  createBlog() {
    const newBlog = {
      title: 'blogger pro max',
      body: 'helloworld123'
    };
    this.AuthService.createContent(newBlog).subscribe({
      next: (res: any) => {
        this.AuthService.setCurrentContent(res);
        this.ngOnInit();
        // this.router.navigate(['/blog-open']);
      },
      error: (err: any) => {
        console.error('Error creating blog:', err);
      }
    });
  }
}
