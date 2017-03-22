import { Component } from '@angular/core'

@Component ({
    selector: 'ez-app',
    templateUrl: './app/entry.component.html'
})
export class EntryComponent {
    
    private greetings: string;

    ngOnInit() {
        //Called after the constructor, initializing input properties, and the first call to ngOnChanges.
        //Add 'implements OnInit' to the class.
        this.greetings = "Welcome to Mirage UI (Entry)";
    }
}
