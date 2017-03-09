package api

import (
	"fmt"
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

func APIAddUser(paras *ApiParas) *ApiResponse {
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

func APILoginByUser(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APILoginByUser\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}

func APIShowUser(paras *ApiParas) *ApiResponse {

	octlog.Debug("running in APIShowUser\n")

	resp := new(ApiResponse)

	userId := paras.InParas.Paras["id"].(string)
	temp := user.FindUser(paras.Db, userId)

	if temp == nil {
		resp.Error = merrors.ERR_SEGMENT_NOT_EXIST
		resp.ErrorLog = fmt.Sprintf("user %s not found", userId)
		return resp
	}

	resp.Data = temp

	octlog.Debug("found User %s", temp.Name)

	return resp
}

func APIUpdateUser(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APIUpdateUser\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}

func APIShowAllUser(paras *ApiParas) *ApiResponse {

	resp := new(ApiResponse)

	offset := ParasInt(paras.InParas.Paras["start"])
	limit := ParasInt(paras.InParas.Paras["limit"])

	rows, err := paras.Db.Query("SELECT ID,U_Name,U_State,U_Type,U_Email,U_PhoneNumber,"+
		"U_Description,U_CreateTime,U_LastLogin,U_LastSync "+
		"FROM tb_user LIMIT ?,?", offset, limit)
	if err != nil {
		logger.Errorf("query user list error %s\n", err.Error())
		resp.Error = merrors.ERR_DB_ERR
		return resp
	}
	defer rows.Close()

	userList := make([]user.User, 0)

	for rows.Next() {
		var user user.User
		err = rows.Scan(&user.Id, &user.Name, &user.State,
			&user.Type, &user.Email, &user.PhoneNumber, &user.Desc,
			&user.CreateTime, &user.LastLogin, &user.LastSync)
		if err == nil {
			logger.Debugf("query result: %s:%d\n", user.Id,
				user.Name, user.State, user.Type)
		} else {
			logger.Errorf("query user list error %s\n", err.Error())
		}
		userList = append(userList, user)
	}

	userCount := user.GetUserCount(paras.Db)

	result := make(map[string]interface{}, 3)
	result["total"] = userCount
	result["count"] = len(userList)
	result["data"] = userList

	resp.Data = result

	return resp
}

func APIDeleteUser(paras *ApiParas) *ApiResponse {

	octlog.Debug("running in APIDeleteUser\n")

	resp := new(ApiResponse)

	user := FindUserById(paras.Db, paras.InParas.Paras["id"].(string))
	if user == nil {
		resp.Error = merrors.ERR_SEGMENT_NOT_EXIST
		return resp
	}

	user.Delete(paras.Db)

	return resp
}

func APIShowUserList(paras *ApiParas) *ApiResponse {

	resp := new(ApiResponse)

	rows, err := paras.Db.Query("SELECT ID,U_Name FROM tb_user")
	if err != nil {
		logger.Errorf("query user list error %s\n", err.Error())
		resp.Error = merrors.ERR_DB_ERR
		return resp
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
		userList = append(userList, user.Brief())
	}

	resp.Data = userList

	return resp
}

func APIResetUserPassword(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APIResetUserPassword\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}

func APIUpdateUserPassword(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APIUpdateUserPassword\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}

func APIUserLogOut(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APILogOut\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}
