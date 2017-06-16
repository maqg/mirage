import { Hero } from './hero';
import { Component, OnInit } from '@angular/core';

@Component ({
    selector: 'sanguo-ui',
    providers: [],
    templateUrl: './app/app.component.html',
    styleUrls: [ './app/app.component.css' ],
})
export class AppComponent implements OnInit {

    private title : string;
    private myHero : Hero;

    private heroes : Hero[];

    private selectedHero: Hero;

    onSelect(hero: Hero) {
        this.selectedHero = hero;
    }

    ngOnInit() {
        this.title = "Heroes In Sanguo";

        this.heroes = [
            new Hero(1, "曹操"),
            new Hero(2, "刘备"),
            new Hero(3, "孙权"),
            new Hero(4, "诸葛亮")
        ];

        this.myHero = this.heroes[0];        
    }
}