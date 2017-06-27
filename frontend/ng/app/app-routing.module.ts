import { AccountAddComponent } from './account-add.component';
import { AccountDetailComponent } from './account-detail.component';
import { ConfigComponent } from './config.component';
import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { DashboardComponent }   from './dashboard.component';
import { AccountsComponent }      from './accounts.component';

const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'dashboard',  component: DashboardComponent },
  { path: 'accounts',     component: AccountsComponent },
  { path: 'accounts/add',     component: AccountAddComponent },
  { path: 'detail/:id', component: AccountDetailComponent },
  { path: 'config',     component: ConfigComponent },
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}