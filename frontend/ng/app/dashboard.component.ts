import { Component, OnInit } from '@angular/core';

import { Account }        from './account';
import { AccountService } from './account.service';

@Component({
  selector: 'my-dashboard',
  templateUrl: './app/dashboard.component.html',
  styleUrls: [ './app/dashboard.component.css' ]
})
export class DashboardComponent implements OnInit {

  accounts: Account[] = [];

  constructor(private accountService: AccountService) { }

  ngOnInit(): void {
    this.accountService.getAccounts()
      .then(accounts => this.accounts = accounts.slice(1, 5));
  }
}
