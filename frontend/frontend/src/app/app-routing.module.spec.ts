import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { Router } from '@angular/router';
import { Location } from '@angular/common';

import { ComponentAComponent } from './component-a/component-a.component';
import { ComponentBComponent } from './component-b/component-b.component';
import { ComponentCComponent } from './component-c/component-c.component';
import { ComponentDComponent } from './component-d/component-d.component';

import { AppRoutingModule } from './app-routing.module';

describe('AppRoutingModule', () => {
  let router: Router;
  let location: Location;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AppRoutingModule,
        RouterTestingModule.withRoutes([])
      ],
      declarations: [
        ComponentAComponent,
        ComponentBComponent,
        ComponentCComponent,
        ComponentDComponent
      ]
    }).compileComponents();

    router = TestBed.inject(Router);
    location = TestBed.inject(Location);
    router.initialNavigation();
  });

  it('navigates to "" -> ComponentAComponent', async () => {
    await router.navigate(['']);
    expect(location.path()).toBe('');
  });

  it('navigates to "user" -> ComponentBComponent', async () => {
    await router.navigate(['user']);
    expect(location.path()).toBe('/user');
  });

  it('navigates to "unauthorized" -> ComponentCComponent', async () => {
    await router.navigate(['unauthorized']);
    expect(location.path()).toBe('/unauthorized');
  });

  it('navigates to "thankyou" -> ComponentDComponent', async () => {
    await router.navigate(['thankyou']);
    expect(location.path()).toBe('/thankyou');
  });

  it('wildcard redirects to ""', async () => {
    await router.navigate(['random']);
    expect(location.path()).toBe('');
  });
});
