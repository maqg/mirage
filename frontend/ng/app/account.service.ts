import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Account } from './account';

@Injectable()
export class AccountService {

  private headers = new Headers({ 'Content-Type': 'application/json' });
  private accountsUrl = 'api/accounts';  // URL to web api

  constructor(private http: Http) { }

  getAccounts(): Promise<Account[]> {
    return this.http.get(this.accountsUrl)
      .toPromise()
      .then(response => response.json().data as Account[])
      .catch(this.handleError);
  }

  getAccount(id: number): Promise<Account> {
    const url = `${this.accountsUrl}/${id}`;
    return this.http.get(url)
      .toPromise()
      .then(response => response.json().data as Account)
      .catch(this.handleError);
  }

  delete(id: number): Promise<void> {
    const url = `${this.accountsUrl}/${id}`;
    return this.http.delete(url, { headers: this.headers })
      .toPromise()
      .then(() => null)
      .catch(this.handleError);
  }

  create(name: string, email: string, phoneNumber: string): Promise<Account> {
    var account = {
      "name": name,
      "email": email,
      "phoneNumber": phoneNumber
    };
    return this.http
      .post(this.accountsUrl, JSON.stringify(account), { headers: this.headers })
      .toPromise()
      .then(res => res.json().data)
      .catch(this.handleError);
  }

  update(account: Account): Promise<Account> {
    const url = `${this.accountsUrl}/${account.id}`;
    return this.http
      .put(url, JSON.stringify(account), { headers: this.headers })
      .toPromise()
      .then(() => account)
      .catch(this.handleError);
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }
}