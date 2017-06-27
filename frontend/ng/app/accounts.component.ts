import { Location } from '@angular/common';
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

    delete(account: Account): void {
        this.accountService
            .delete(account.id)
            .then(() => {
                this.accounts = this.accounts.filter(h => h !== account);
            });
    }

    gotoDetail(account: Account): void {
        this.router.navigate(['/detail', account.id]);
    }

    gotoAdd(): void {
        this.router.navigate(['/accounts/add']);
    }
}