package api

import (
	"octlink/mirage/src/modules/user"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
	"octlink/mirage/src/utils/uuid"
)

func FindUserByName(db *octmysql.OctMysql, name string) *user.User {

	row := db.QueryRow("SELECT ID,U_Name FROM tb_user WHERE U_Name = ? LIMIT 1", name)

	user := new(user.User)

	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		logger.Errorf("Find user %s error %s", name, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", user.Id, user.Name)

	return user
}

func FindUserById(db *octmysql.OctMysql, id string) *user.User {

	row := db.QueryRow("SELECT ID,U_Name FROM tb_user WHERE ID = ? LIMIT 1", id)

	user := new(user.User)
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		logger.Errorf("Find user %s error %s", id, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", user.Id, user.Name)

	return user
}

func APIAddAccount(paras *ApiParas) *ApiResponse {
	resp := new(ApiResponse)

	newUser := FindUserByName(paras.Db, paras.InParas.Paras["account"].(string))
	if newUser != nil {
		logger.Errorf("user %s already exist\n", newUser.Name)
		resp.Error = merrors.ERR_SEGMENT_ALREADY_EXIST
		return resp
	}

	newUser = new(user.User)
	newUser.Id = uuid.Generate().Simple()
	newUser.Name = paras.InParas.Paras["account"].(string)
	newUser.Type = 0
	newUser.Email = paras.InParas.Paras["email"].(string)
	newUser.PhoneNumber = paras.InParas.Paras["phoneNumber"].(string)

	resp.Error = newUser.Add(paras.Db)

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

	rows, err := paras.Db.Query("SELECT ID,U_Name FROM tb_user")
	if err != nil {
		logger.Errorf("query user list error %s\n", err.Error())
	}
	defer rows.Close()

	userList := make([]map[string]interface{}, 0)

	for rows.Next() {
		var user user.User
		err = rows.Scan(&user.Id, &user.Name)
		if err == nil {
			logger.Debugf("query result: %s:%s\n", user.Id, user.Name)
		} else {
			logger.Errorf("query user list error %s\n", err.Error())
		}
		userList = append(userList, user.UserBrief())
	}

	resp.Data = userList
	resp.Error = 0
	return resp
}

func APIDeleteAccount(paras *ApiParas) *ApiResponse {

	octlog.Debug("running in APIDeleteAccount\n")

	resp := new(ApiResponse)

	user := FindUserById(paras.Db, paras.InParas.Paras["id"].(string))
	if user == nil {
		resp.Error = merrors.ERR_SEGMENT_NOT_EXIST
		return resp
	}

	user.Delete(paras.Db)

	return resp
}

func APIShowAllAccount(paras *ApiParas) *ApiResponse {
	resp := new(ApiResponse)

	octlog.Debug("running in APIShowAllAccount\n")

	rows, err := paras.Db.Query("SELECT ID,U_Name,U_State,U_Type FROM tb_user")
	if err != nil {
		logger.Errorf("query user list error %s\n", err.Error())
		resp.Error = merrors.ERR_DB_ERR
		return resp
	}
	defer rows.Close()

	userList := make([]user.User, 0)

	for rows.Next() {
		var user user.User
		err = rows.Scan(&user.Id, &user.Name, &user.State, &user.Type)
		if err == nil {
			logger.Debugf("query result: %s:%s\n", user.Id, user.Name, user.State, user.Type)
		} else {
			logger.Errorf("query user list error %s\n", err.Error())
		}
		userList = append(userList, user)
	}

	resp.Data = userList
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
