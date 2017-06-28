export class ApiRequest {
    module: string;
    api: string;
    paras: {};
    async: boolean;
    session: {
        'uuid': string,
        'skey': string
    };
}

export class ApiResponse {
    errorObj: {
        errorNo: number,
        errorMsg: string,
        errorLog: string,
        errorMsgEN: string
    };
    apiId: string;
    session: {
        uuid: string
    };
    data: any;
}

export class ServiceResponse {
    errorCode: number;
    errorMsg: string;
    api: string;
    data: any;
}

export const API_PATH = '/api/';
export const API_PREFIX = "octlink.mirage.center";
export const API_KEY = "00000000000000000000000000000000";
export const SESSION_ID = "00000000000000000000000000000000";