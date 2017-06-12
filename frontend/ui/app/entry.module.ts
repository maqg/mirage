
import { NgModule } from '@angular/core'
import { BrowserModule } from '@angular/platform-browser'
import { FormsModule } from '@angular/forms'

import { EntryComponent } from './entry.component'

@NgModule({
  imports: [ BrowserModule, FormsModule ],
  declarations: [ EntryComponent ],
  bootstrap: [ EntryComponent ],
})
export class EntryModule { }