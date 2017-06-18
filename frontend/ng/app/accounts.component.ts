import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Account } from './account';
import { AccountService } from './account.service';

@Component({
    selector: 'my-accounts',
    templateUrl: './app/accounts.component.html',
    styleUrls: ['./app/accounts.component.css']
})

export class AccountsComponent implements OnInit {
    accounts: Account[];
    selectAccount: Account;

    constructor(
        private accountService: AccountService,
        private router: Router) { }

    getAccounts(): void {
        this.accountService
            .getAccounts()
            .then(accounts => this.accounts = accounts);
    }

    ngOnInit(): void {
        this.getAccounts();
    }

    onSelect(account: Account): void {
        this.selectAccount = account;
    }
}