
import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'mirage-ui',
    providers: [],
    styleUrls: [ './app/app.component.css' ],
    templateUrl: 'app/app.component.html',
})
export class AppComponent implements OnInit {
    private where : string;

    ngOnInit() {
        this.where = "China";
    }
}