import { Config } from './config';
import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'my-config',
    templateUrl: './app/config.component.html',
    styleUrls: ['./app/config.component.css']
})
export class ConfigComponent implements OnInit {

    private config: Config;

    ngOnInit(): void {
        this.config = {
            version: "0.0.1",
            name: "Mirage"
        }
    }
}
