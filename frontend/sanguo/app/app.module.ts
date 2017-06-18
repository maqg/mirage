import { HeroService } from './hero.service';
import { DashboardComponent } from './dashboard.component';
import { HeroesComponent } from './heroes.component';
import { AppComponent } from './app.component';
import { HeroDetailComponent } from './hero-detail.component';

import { NgModule } from '@angular/core';
import { HttpModule } from '@angular/http';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';

// Imports for loading & configuring the in-memory web api
import { InMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InMemoryDataService }  from './in-memory-data.service';


@NgModule({
    imports: [
        HttpModule,
        BrowserModule,
        FormsModule,
        AppRoutingModule,

        InMemoryWebApiModule.forRoot(InMemoryDataService),        
    ],

    declarations: [
        AppComponent,
        HeroDetailComponent,
        HeroesComponent,
        DashboardComponent
    ],

    providers: [HeroService],

    bootstrap: [
        AppComponent,
    ]
})
export class AppModule {
}