import { InMemoryDbService } from 'angular-in-memory-web-api';

export class InMemoryDataService implements InMemoryDbService {
    createDb() {
        let heroes = [
            { id: 1, name: '曹操' },
            { id: 11, name: '刘备' },
            { id: 20, name: '孙权' },
            { id: 35, name: '诸葛亮' },
            { id: 40, name: '张飞' },
            { id: 41, name: '赵云' },
            { id: 42, name: '关羽' },
            { id: 43, name: '曹仁' },
            { id: 44, name: '李典' },
            { id: 46, name: '马超' },
            { id: 48, name: '张辽' },
        ];

        return {
            heroes
        };
    }
}