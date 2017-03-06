package api

import (
	"octlink/mirage/src/modules/account"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
	"octlink/mirage/src/utils/uuid"
)

func FindAccountByName(db *octmysql.OctMysql, name string) *account.Account {

	row := db.QueryRow("SELECT ID,U_Name FROM tb_account WHERE U_Name = ? LIMIT 1", name)

	account := new(account.Account)

	err := row.Scan(&account.Id, &account.Name)
	if err != nil {
		logger.Errorf("Find account %s error %s", name, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", account.Id, account.Name)

	return account
}

func FindAccountById(db *octmysql.OctMysql, id string) *account.Account {

	row := db.QueryRow("SELECT ID,U_Name FROM tb_account WHERE ID = ? LIMIT 1", id)

	account := new(account.Account)
	err := row.Scan(&account.Id, &account.Name)
	if err != nil {
		logger.Errorf("Find account %s error %s", id, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", account.Id, account.Name)

	return account
}

func APIAddAccount(paras *ApiParas) *ApiResponse {
	resp := new(ApiResponse)

	newAccount := FindAccountByName(paras.Db, paras.InParas.Paras["account"].(string))
	if newAccount != nil {
		logger.Errorf("account %s already exist\n", newAccount.Name)
		resp.Error = merrors.ERR_SEGMENT_ALREADY_EXIST
		return resp
	}

	newAccount = new(account.Account)
	newAccount.Id = uuid.Generate().Simple()
	newAccount.Name = paras.InParas.Paras["account"].(string)
	newAccount.Type = 0
	newAccount.Email = paras.InParas.Paras["email"].(string)
	newAccount.PhoneNumber = paras.InParas.Paras["phoneNumber"].(string)

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
	resp.Error = 0
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
	resp.Error = 0
	return resp
}

func APIDeleteAccount(paras *ApiParas) *ApiResponse {

	octlog.Debug("running in APIDeleteAccount\n")

	resp := new(ApiResponse)

	account := FindAccountById(paras.Db, paras.InParas.Paras["id"].(string))
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

	rows, err := paras.Db.Query("SELECT ID,U_Name,U_State,U_Type FROM tb_account")
	if err != nil {
		logger.Errorf("query account list error %s\n", err.Error())
		resp.Error = merrors.ERR_DB_ERR
		return resp
	}
	defer rows.Close()

	accountList := make([]account.Account, 0)

	for rows.Next() {
		var account account.Account
		err = rows.Scan(&account.Id, &account.Name, &account.State, &account.Type)
		if err == nil {
			logger.Debugf("query result: %s:%s\n", account.Id, account.Name, account.State, account.Type)
		} else {
			logger.Errorf("query account list error %s\n", err.Error())
		}
		accountList = append(accountList, account)
	}

	resp.Data = accountList
	resp.Error = 0
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
