import { Component } from '@angular/core';

import { Hero } from './hero';
import { HeroService } from './hero.service';



@Component({
  selector: 'mirage-ui',
  providers: [ HeroService ],
  templateUrl: './hero/app.component.html',
  styleUrls: [ './hero/app.component.css' ]
})
export class AppComponent {
  title = 'Tour of Heroes';
}