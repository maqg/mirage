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

  private createGetAccountReq(accountId: string) {
    let req = new ApiRequest();
    req.module = "account";
    req.api = API_PREFIX + "." + req.module + "." + "APIShowAccount";
    req.paras = {
      id: accountId
    };
    req.async = false;
    req.session = {
      uuid: SESSION_ID,
      skey: API_KEY
    };
    return req;
  }

  private createDeleteAccountReq(accountId: string) {
    let req = new ApiRequest();
    req.module = "account";
    req.api = API_PREFIX + "." + req.module + "." + "APIDeleteAccount";
    req.paras = {
      id: accountId
    };
    req.async = false;
    req.session = {
      uuid: SESSION_ID,
      skey: API_KEY
    };
    return req;
  }

  private createUpdateAccountReq(accountId: string, email: string, phoneNumber: string, desc: string) {
    let req = new ApiRequest();
    req.module = "account";
    req.api = API_PREFIX + "." + req.module + "." + "APIUpdateAccount";
    req.paras = {
      id: accountId,
      email: email,
      phoneNumber: phoneNumber,
      desc: desc
    };
    req.async = false;
    req.session = {
      uuid: SESSION_ID,
      skey: API_KEY
    };
    return req;
  }

  private createAddAccountReq(name: string, password: string, email: string, phoneNumber: string, desc: string) {
    let req = new ApiRequest();
    req.module = "account";
    req.api = API_PREFIX + "." + req.module + "." + "APIAddAccount";
    req.paras = {
      account: name,
      password: password,
      email: email,
      type: 3,
      phoneNumber: phoneNumber,
      desc: desc
    };
    req.async = false;
    req.session = {
      uuid: SESSION_ID,
      skey: API_KEY
    };
    return req;
  }

  getAccounts(): Promise<Account[]> {

    let req = this.createGetAccountsReq();
    return this.apiService.doRequest(req)
      .then(res => {
        if (res.errorCode !== 0) {
          return [];
        }
        let objs = res.data.data;
        var accounts: Account[] = [];
        for (let obj of objs) {
          let account = obj as Account;
          accounts.push(account);
        }
        return accounts;
      })
      .catch(err => {
        return Promise.reject(err);
      });
  }

  getAccount(id: string): Promise<Account> {
    let req = this.createGetAccountReq(id);
    return this.apiService.doRequest(req)
      .then(res => {
        if (res.errorCode !== 0) {
          return null;
        }
        return res.data as Account;
      })
      .catch(err => {
        return Promise.reject(err);
      });
  }

  delete(id: string): Promise<void> {
    let req = this.createDeleteAccountReq(id);
    return this.apiService.doRequest(req)
      .then(res => {
        if (res.errorCode !== 0) {
          return;
        }
      })
      .catch(err => {
        this.handleError;
      });
  }

  create(name: string, password: string, email: string, phoneNumber: string): Promise<void> {
    let req = this.createAddAccountReq(name, password, email, phoneNumber, "");
    return this.apiService.doRequest(req)
      .then(res => {
        if (res.errorCode !== 0) {
          return;
        }
      })
      .catch(err => {
        this.handleError;
      });
  }

  update(account: Account): Promise<Account> {
    let req = this.createUpdateAccountReq(account.id, account.email, account.phoneNumber, account.desc);
    return this.apiService.doRequest(req)
      .then(res => {
        if (res.errorCode !== 0) {
          return null;
        }
      })
      .catch(err => {
        this.handleError;
      });
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }
}