import 'reflect-metadata';
import 'zone.js';

import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { EntryModule } from './app/entry.module'

platformBrowserDynamic().bootstrapModule(EntryModule);