import 'reflect-metadata'
import 'zone.js'

import { platformBrowserDynamic } from '@angular/platform-browser-dynamic'

import { AppModule } from './app/app.module'

platformBrowserDynamic().bootstrapModule(AppModule)
    .then(success => console.log('BootStrap success'))
    .catch(error => console.log(error))