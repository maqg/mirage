import { ApiRequest, ApiResponse, API_PATH, ServiceResponse } from './api';
import { Router } from '@angular/router';
import { Injectable } from '@angular/core';
import { Headers, Http, Response } from '@angular/http';
import { TimeoutError } from 'rxjs';
import 'rxjs/add/operator/toPromise';

@Injectable()
export class ApiService {

    constructor(private _http: Http, private _router: Router) {
    }

    doRequest(req: ApiRequest): Promise<ServiceResponse> {
        return this._http.post(API_PATH, req, null)
            .toPromise()
            .then(res => this._getResponse(res, req))
            .catch(er => {
                return this._handleError(er, req);
            });
    }

    private _getResponse(res: Response, req: ApiRequest) {
        
        let respObj: ApiResponse = res.json() as ApiResponse;
        var serviceResponse = new ServiceResponse();
        serviceResponse.api = req.api;
        serviceResponse.data = respObj.data;
        serviceResponse.errorCode = respObj.errorObj.errorNo;
        serviceResponse.errorMsg = respObj.errorObj.errorMsg;

        return serviceResponse;
    }

    private _handleError(error: any, req: ApiRequest) {
        var resp: ServiceResponse = new ServiceResponse();
        resp.api = req.api;
        resp.errorCode = 100;
        resp.errorMsg = "API Request Error";
        return resp;
    }
}