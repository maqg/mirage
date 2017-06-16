import { HeroService } from './hero.service';
import { Hero } from './hero';
import { Component, OnInit } from '@angular/core';

@Component ({
    selector: 'sanguo-ui',
    providers: [
        HeroService
    ],
    templateUrl: './app/app.component.html',
    styleUrls: [
        './app/app.component.css'
    ],
})
export class AppComponent implements OnInit {

    private title : string;

    private heroes : Hero[];

    private selectedHero: Hero;

    constructor(private heroService: HeroService) {
    }

    onSelect(hero: Hero) {
        this.selectedHero = hero;
    }

    getHeroes() {
        this.heroService.getHeroes().then(heroes=>this.heroes=heroes);
    }

    ngOnInit() {
        this.title = "三国英雄榜";
        this.getHeroes();
    }
}