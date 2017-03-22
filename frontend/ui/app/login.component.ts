
import { Component } from '@angular/core'

@Component ({
    selector: 'mirage-ui',
    templateUrl: './app/login.component.html',
    styleUrls: [ './css/login.css' ]
})
export class LoginComponent {
    
    private greetings: string;

    ngOnInit() {
        //Called after the constructor, initializing input properties, and the first call to ngOnChanges.
        //Add 'implements OnInit' to the class.
        this.greetings = "Welcome to Mirage UI";
    }
}