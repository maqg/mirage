import { HeroService } from './hero.service';
import { Hero } from './hero';
import { Component, Input, OnInit } from '@angular/core';

import { ActivatedRoute, Params } from '@angular/router';
import { Location } from '@angular/common';
import 'rxjs/add/operator/switchMap';

@Component({
    moduleId: '1',
    selector: 'my-hero-detail',
    providers: [
        HeroService
    ],
    styleUrls: [
        './app/hero-detail.component.css'
    ],
    templateUrl: './app/hero-detail.component.html',
})

export class HeroDetailComponent implements OnInit {
    @Input()
    hero: Hero;

    ngOnInit(): void {

        this.route.params
            .switchMap((params: Params) => this.heroService.getHero(+params['id']))
            .subscribe(hero => this.hero = hero);

    }

    goBack(): void {
        this.location.back();
    }

    save(): void {
        this.heroService.update(this.hero)
            .then(() => this.goBack());
    }

    constructor(
        private heroService: HeroService,
        private route: ActivatedRoute,
        private location: Location
    ) { }
}