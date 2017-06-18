import { AppComponent } from './app.component';

import { NgModule } from '@angular/core'

import { BrowserModule } from '@angular/platform-browser'

import { InMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InMemoryDataService } from './in-memory-data.service';

@NgModule({
    imports: [
        BrowserModule,
        InMemoryWebApiModule.forRoot(InMemoryDataService),
    ],

    declarations: [
        AppComponent
    ],

    providers: [],
    bootstrap: [
        AppComponent
    ],
})
export class AppModule { }
