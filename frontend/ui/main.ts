import 'reflect-metadata';
import 'zone.js';

import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { LoginModule } from './app/login.module'

platformBrowserDynamic().bootstrapModule(LoginModule);