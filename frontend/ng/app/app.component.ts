
import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'mirage-ui',
    providers: [],
    templateUrl: 'app/app.component.html',
})
export class AppComponent implements OnInit {
    private where : string;

    ngOnInit() {
        this.where = "China";
    }
}