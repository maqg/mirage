import { InMemoryDbService } from 'angular-in-memory-web-api';

export class InMemoryDataService implements InMemoryDbService {
  createDb() {
    let accounts = [
      {
        id: 11,
        name: 'admin',
        email: "admin@octink.com",
        phoneNumber: "12911342134",
        lastSync: "2017-12-12 22:33:44",
        createTime: "2017-12-12 22:33:44",
        lastLogin: "2017-12-12 22:33:44",
      },
      {id: 12, name: 'Narco'},
      {id: 13, name: 'Bombasto'},
      {id: 14, name: 'Celeritas'},
      {id: 15, name: 'Magneta'},
      {id: 16, name: 'RubberMan'},
      {id: 17, name: 'Dynama'},
      {id: 18, name: 'Dr IQ'},
      {id: 19, name: 'Magma'},
      {id: 20, name: 'Tornado'}
    ];
    return {
        accounts
    };
  }
}
