import { Hero } from './hero';
import { Component, Input } from '@angular/core';


@Component({
    selector: 'my-hero-detail',
    templateUrl: './app/hero-detail.component.html',
})

export class HeroDetailComponent {
    @Input()
    hero: Hero;
}