import { ApiService } from './common/api.service';
import { ApiRequest, SESSION_ID, API_KEY, API_PREFIX } from './common/api';
import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Account } from './account';

@Injectable()
export class AccountService {

  private headers = new Headers({ 'Content-Type': 'application/json' });
  private accountsUrl = 'apis/accounts';  // URL to web api

  constructor(private http: Http, private apiService: ApiService) { }

  private createGetAccountsReq() {
    let req = new ApiRequest();
    req.module = "account";
    req.api = API_PREFIX + "." + req.module + "." + "APIShowAllAccount";
    req.paras = {
      start: 0,
      limit: 15
    };
    req.async = false;
    req.session = {
      uuid: SESSION_ID,
      skey: API_KEY
    };
    return req;
  }

  /*
  getAccounts(): Promise<Account[]> {
    return this.http.get(this.accountsUrl)
      .toPromise()
      .then(response => response.json().data as Account[])
      .catch(this.handleError);
  }*/

  getAccounts(): Promise<Account[]> {

    let req = this.createGetAccountsReq();
    return this.apiService.doRequest(req)
      .then(res => {
        if (res.errorCode !== 0) {
          return res.errorMsg;
        }
        var accounts: Account[] = res.data.data.json() as Account[];
        return accounts;
      })
      .catch(err => {
        return Promise.reject(err);
      });
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