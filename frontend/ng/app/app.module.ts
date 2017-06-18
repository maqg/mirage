import { HttpModule } from '@angular/http';
import { AccountService } from './account.service';
import { AccountsComponent } from './accounts.component';
import { DashboardComponent } from './dashboard.component';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgModule } from '@angular/core'
import { FormsModule }   from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser'

import { InMemoryWebApiModule } from 'angular-in-memory-web-api';
import { InMemoryDataService } from './in-memory-data.service';

@NgModule({
    imports: [
        HttpModule,
        FormsModule,
        BrowserModule,
        InMemoryWebApiModule.forRoot(InMemoryDataService),
        AppRoutingModule,
    ],

    declarations: [
        AppComponent,
        DashboardComponent,
        AccountsComponent,
    ],

    providers: [
        AccountService,
    ],
    bootstrap: [
        AppComponent
    ],
})
export class AppModule { }
