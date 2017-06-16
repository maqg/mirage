import { Component, OnInit } from '@angular/core';

@Component ({
    selector: 'sanguo-ui',
    templateUrl: './app/app.component.html',
    styleUrls: [
        './app/app.component.css'
    ],
})
export class AppComponent implements OnInit {

    private title : string;

    ngOnInit() {
        this.title = "三国英雄榜";
    }
}