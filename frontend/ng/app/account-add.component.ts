import { Component } from '@angular/core';
import { Location } from '@angular/common';
import { AccountService } from './account.service';
import { Account } from './account';

@Component({
    selector: 'account-add',
    styleUrls: ['./app/account-add.component.css'],
    templateUrl: './app/account-add.component.html',
})
export class AccountAddComponent {

    account: Account;

    constructor(
        private accountService: AccountService,
        private location: Location
    ) { }

    add(name: string, email: string, phoneNumber: string): void {
        name = name.trim();
        if (!name) {
            this.location.back();
            return;
        }
        this.accountService.create(name, email, phoneNumber);
        this.location.back();
    }

    goBack(): void {
        this.location.back();
    }
}