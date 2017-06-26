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
    selectedAccount: Account;

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
        this.selectedAccount = account;
    }

    delete(account: Account): void {
        this.accountService
            .delete(account.id)
            .then(() => {
                this.accounts = this.accounts.filter(h => h !== account);
                if (this.selectedAccount === account) { this.selectedAccount = null; }
            });
    }

    gotoDetail(): void {
        this.router.navigate(['/detail', this.selectedAccount.id]);
    }
}