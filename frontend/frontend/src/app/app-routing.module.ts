import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ComponentAComponent } from './component-a/component-a.component';
import { ComponentBComponent } from './component-b/component-b.component';
import { ComponentCComponent } from './component-c/component-c.component';

const routes: Routes = [
  { path: '', component: ComponentAComponent },
  { path: 'user', component: ComponentBComponent },
  { path: 'unauthorized', component: ComponentCComponent },
  { path: '**', redirectTo: '' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
