import { Account } from './account';

import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'mirage-ui',
    providers: [],
    styleUrls: [ './app/app.component.css' ],
    templateUrl: 'app/app.component.html',
})
export class AppComponent implements OnInit {
    private where : string;

    private accounts: Account[];

    ngOnInit() {
        this.where = "China";

        this.accounts = [
            { id: 1, name: "admin" },
            { id: 2, name: "root" },
            { id: 3, name: "henry" },
            { id: 4, name: "jacky" },
            { id: 5, name: "bruce" },
        ];

    }
}