import { HeroService } from './hero.service';
import { Component, OnInit } from '@angular/core';
import { Hero } from './hero';

@Component({
    moduleId: '1',
    selector: 'my-dashboard',
    styleUrls: [
        './app/dashboard.component.css'
    ],
    templateUrl: './app/dashboard.component.html',
})
export class DashboardComponent implements OnInit {

    heroes: Hero[] = [];

    constructor(private heroService: HeroService) {
    }

    ngOnInit(): void {
        this.heroService.getHeroes().then(heroes => this.heroes = heroes.slice(1, 9));
    }

}