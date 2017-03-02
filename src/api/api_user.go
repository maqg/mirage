package api

import (
	"fmt"
	"octlink/mirage/src/utils/config"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
	"octlink/mirage/src/utils/uuid"
)

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	State       int    `json:"state"`
	Type        int    `json:"type"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func (user *User) UserBrief() map[string]interface{} {
	u := make(map[string]interface{}, 2)
	u["id"] = user.Id
	u["name"] = user.Name

	return u
}

func (user *User) Add(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("INSERT INTO %s (ID, U_Name, U_Type) VALUES ('%s', '%s', '%d')",
		config.TB_USER, user.Id, user.Name, user.Type)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("insert db error %s", sql)
		return merrors.ERR_DB_ERR
	}

	return 0
}

func (user *User) Delete(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("DELETE FROM %s WHERE ID='%s'", config.TB_USER, user.Id)
	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("delete user %s error", user.Id)
		return merrors.ERR_DB_ERR
	}

	octlog.Debug(sql)

	return 0
}

func FindUserByName(db *octmysql.OctMysql, name string) *User {

	row := db.QueryRow("SELECT ID,U_Name FROM tb_user WHERE U_Name = ? LIMIT 1", name)

	user := new(User)

	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		logger.Errorf("Find user %s error %s", name, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", user.Id, user.Name)

	return user
}

func FindUserById(db *octmysql.OctMysql, id string) *User {

	row := db.QueryRow("SELECT ID,U_Name FROM tb_user WHERE ID = ? LIMIT 1", id)

	user := new(User)
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

	user := FindUserByName(paras.Db, paras.InParas.Paras["account"].(string))
	if user != nil {
		logger.Errorf("user %s already exist\n", user.Name)
		resp.Error = merrors.ERR_SEGMENT_ALREADY_EXIST
		return resp
	}

	user = new(User)
	user.Id = uuid.Generate().Simple()
	user.Name = paras.InParas.Paras["account"].(string)
	user.Type = 0
	user.Email = paras.InParas.Paras["email"].(string)
	user.PhoneNumber = paras.InParas.Paras["phoneNumber"].(string)

	resp.Error = user.Add(paras.Db)

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
		var user User
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

	userList := make([]User, 0)

	for rows.Next() {
		var user User
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
