import { AppComponent } from './app.component';

import { NgModule } from '@angular/core'

import { BrowserModule } from '@angular/platform-browser'

@NgModule({
    imports: [
        BrowserModule,
    ],

    declarations: [
        AppComponent
    ],
    providers: [],
    bootstrap: [
        AppComponent
    ],
})
export class AppModule {}
