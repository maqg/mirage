import 'reflect-metadata';
import 'zone.js';

import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { AppModule } from './hero/app.module'

platformBrowserDynamic().bootstrapModule(AppModule);