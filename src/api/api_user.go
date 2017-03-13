package api

import (
	"fmt"
	"octlink/mirage/src/modules/user"
	"octlink/mirage/src/modules/usergroup"
	"octlink/mirage/src/utils"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/uuid"
)

func APIAddUser(paras *ApiParas) *ApiResponse {
	resp := new(ApiResponse)

	newUser := user.FindUserByName(paras.Db, paras.InParas.Paras["account"].(string))
	if newUser != nil {
		logger.Errorf("user %s already exist\n", newUser.Name)
		resp.Error = merrors.ERR_SEGMENT_ALREADY_EXIST
		return resp
	}

	gid := paras.InParas.Paras["groupId"].(string)
	group := usergroup.FindGroup(paras.Db, gid)
	if group == nil {
		resp.Error = merrors.ERR_USERGROUP_NOT_EXIST
		return resp
	}

	newUser = new(user.User)
	newUser.Id = uuid.Generate().Simple()
	newUser.Name = paras.InParas.Paras["account"].(string)
	newUser.Type = 0
	newUser.Email = paras.InParas.Paras["email"].(string)
	newUser.PhoneNumber = paras.InParas.Paras["phoneNumber"].(string)

	resp.Error = newUser.Add(paras.Db)
	if resp.Error == 0 {
		group.AllowUser(paras.Db, newUser.Id)
	}

	return resp
}

func APILoginByUser(paras *ApiParas) *ApiResponse {

	resp := new(ApiResponse)

	uid := paras.InParas.Paras["account"].(string)
	password := paras.InParas.Paras["password"].(string)

	logger.Debugf("Login %s:%s", uid, password)

	user := user.FindUserByName(paras.Db, uid)
	if user == nil {
		logger.Errorf("account %s already exist\n", uid)
		resp.Error = merrors.ERR_USER_NOT_EXIST
		return resp
	}

	session := user.Login(paras.Db, password)
	if session == nil {
		resp.Error = merrors.ERR_PASSWORD_DONT_MATCH
		return resp
	}

	resp.Data = session

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

	offset := utils.ParasInt(paras.InParas.Paras["start"])
	limit := utils.ParasInt(paras.InParas.Paras["limit"])

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

	user := user.FindUser(paras.Db, paras.InParas.Paras["id"].(string))
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

	userList := make([]map[string]string, 0)

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
