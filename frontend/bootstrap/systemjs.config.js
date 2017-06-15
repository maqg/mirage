// map tells the System loader where to look for things
var map = {
 'app':      'dist', // 'dist',
 '@angular':     'node_modules/@angular',
 'angular2-in-memory-web-api': 'node_modules/angular2-in-memory-web-api',
 'rxjs':      'node_modules/rxjs',
 // add ng-bootstrap location map 
 '@ng-bootstrap':    'node_modules/@ng-bootstrap'
};
// packages tells the System loader how to load when no filename and/or no extension
var packages = {
 'app': { main: 'main.js', defaultExtension: 'js', format: 'amd' },
 'rxjs': { defaultExtension: 'js' },
 'angular2-in-memory-web-api': { main: 'index.js', defaultExtension: 'js' },
 // add ng-bootstrap package config
 '@ng-bootstrap/ng-bootstrap': { main: 'index.js', defaultExtension: 'js' }
};