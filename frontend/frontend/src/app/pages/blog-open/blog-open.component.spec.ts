import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BlogOpenComponent } from './blog-open.component';

describe('BlogOpenComponent', () => {
  let component: BlogOpenComponent;
  let fixture: ComponentFixture<BlogOpenComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BlogOpenComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BlogOpenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
