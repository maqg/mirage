package api

import (
	"fmt"
	"octlink/mirage/src/modules/account"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/uuid"
)

func ParasInt(val interface{}) int {
	return int(val.(float64))
}

func APIAddAccount(paras *ApiParas) *ApiResponse {
	resp := new(ApiResponse)

	newAccount := account.FindAccountByName(paras.Db, paras.InParas.Paras["account"].(string))
	if newAccount != nil {
		logger.Errorf("account %s already exist\n", newAccount.Name)
		resp.Error = merrors.ERR_SEGMENT_ALREADY_EXIST
		return resp
	}

	newAccount = new(account.Account)
	newAccount.Id = uuid.Generate().Simple()
	newAccount.Name = paras.InParas.Paras["account"].(string)
	newAccount.Type = ParasInt(paras.InParas.Paras["type"])
	newAccount.Email = paras.InParas.Paras["email"].(string)
	newAccount.PhoneNumber = paras.InParas.Paras["phoneNumber"].(string)
	newAccount.Password = paras.InParas.Paras["password"].(string)
	newAccount.Desc = paras.InParas.Paras["desc"].(string)

	resp.Error = newAccount.Add(paras.Db)

	return resp
}

func APILoginByAccount(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APILoginByAccount\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}

func APIShowAccount(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APIShowAccount\n")
	resp := new(ApiResponse)

	accountId := paras.InParas.Paras["id"].(string)
	temp := account.FindAccount(paras.Db, accountId)

	if temp == nil {
		resp.Error = merrors.ERR_SEGMENT_NOT_EXIST
		resp.ErrorLog = fmt.Sprintf("user %s not found", accountId)
		return resp
	}

	resp.Data = temp

	octlog.Debug("found User %s", temp.Name)

	return resp
}

func APIUpdateAccount(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APIUpdateAccount\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}

func APIShowAccountList(paras *ApiParas) *ApiResponse {

	resp := new(ApiResponse)

	octlog.Debug("running in APIShowAccountList\n")

	rows, err := paras.Db.Query("SELECT ID,U_Name FROM tb_account")
	if err != nil {
		logger.Errorf("query account list error %s\n", err.Error())
	}
	defer rows.Close()

	accountList := make([]map[string]interface{}, 0)

	for rows.Next() {
		var account account.Account
		err = rows.Scan(&account.Id, &account.Name)
		if err == nil {
			logger.Debugf("query result: %s:%s\n", account.Id, account.Name)
		} else {
			logger.Errorf("query account list error %s\n", err.Error())
		}
		accountList = append(accountList, account.Brief())
	}

	resp.Data = accountList

	return resp
}

func APIDeleteAccount(paras *ApiParas) *ApiResponse {

	octlog.Debug("running in APIDeleteAccount\n")

	resp := new(ApiResponse)

	account := account.FindAccount(paras.Db,
		paras.InParas.Paras["id"].(string))
	if account == nil {
		resp.Error = merrors.ERR_SEGMENT_NOT_EXIST
		return resp
	}

	account.Delete(paras.Db)

	return resp
}

func APIShowAllAccount(paras *ApiParas) *ApiResponse {
	resp := new(ApiResponse)

	octlog.Debug("running in APIShowAllAccount\n")

	offset := ParasInt(paras.InParas.Paras["start"])
	limit := ParasInt(paras.InParas.Paras["limit"])

	rows, err := paras.Db.Query("SELECT ID,U_Name,U_State,U_Type,U_Email,U_PhoneNumber,"+
		"U_Description,U_CreateTime,U_LastLogin,U_LastSync "+
		"FROM tb_account LIMIT ?,?", offset, limit)
	if err != nil {
		logger.Errorf("query account list error %s\n", err.Error())
		resp.Error = merrors.ERR_DB_ERR
		return resp
	}
	defer rows.Close()

	accountList := make([]account.Account, 0)

	for rows.Next() {
		var account account.Account
		err = rows.Scan(&account.Id, &account.Name, &account.State,
			&account.Type, &account.Email, &account.PhoneNumber, &account.Desc,
			&account.CreateTime, &account.LastLogin, &account.LastSync)
		if err == nil {
			logger.Debugf("query result: %s:%s\n", account.Id,
				account.Name, account.State, account.Type)
		} else {
			logger.Errorf("query account list error %s\n", err.Error())
		}
		accountList = append(accountList, account)
	}

	count := account.GetAccountCount(paras.Db)

	result := make(map[string]interface{}, 3)
	result["total"] = count
	result["count"] = len(accountList)
	result["data"] = accountList

	resp.Data = result

	return resp
}

func APIResetAccountPassword(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APIResetAccountPassword\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}

func APIUpdateAccountPassword(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APIUpdateAccountPassword\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}

func APILogOut(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APILogOut\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}
