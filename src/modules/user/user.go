package user

import (
	"fmt"
	"octlink/mirage/src/utils/config"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
)

var logger *octlog.LogConfig

func InitLog(level int) {
	logger = octlog.InitLogConfig("user.log", level)
}

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	State       int    `json:"state"` // 1 for enable,0 for disabled
	Type        int    `json:"type"`  // 1 for terminal user
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"passord"`
	LastLogin   int64  `json:"lastLogin"`
	CreateTime  int64  `json:"createTime"`
	LastSync    int64  `json:"lastSync"`
	Desc        string `json:"desc"`
}

func (user *User) Brief() map[string]string {
	return map[string]string{
		"id":   user.Id,
		"name": user.Name,
	}
}

func GetUserCount(db *octmysql.OctMysql) int {
	count, _ := db.Count(config.TB_USER, "")
	return count
}

func (user *User) Add(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("INSERT INTO %s (ID, U_Name, U_Type, "+
		"U_Email, U_PhoneNumber, U_Password, U_CreateTime, "+
		"U_Description) VALUES ('%s', '%s', '%d', '%s', '%s', "+
		"'%s', '%d', '%s')",
		config.TB_USER,
		user.Id, user.Name, user.Type,
		user.Email, user.PhoneNumber, user.Password,
		user.CreateTime, user.Desc)

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

	row := db.QueryRow("SELECT ID,U_Name,U_State,U_Type,U_Email,U_PhoneNumber,"+
		"U_Description,U_CreateTime,U_LastLogin,U_LastSync "+
		"FROM tb_user WHERE U_Name = ? LIMIT 1", name)

	user := new(User)

	err := row.Scan(&user.Id, &user.Name, &user.State,
		&user.Type, &user.Email, &user.PhoneNumber, &user.Desc,
		&user.CreateTime, &user.LastLogin, &user.LastSync)
	if err != nil {
		logger.Errorf("Find account %s error %s", name, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", user.Id, user.Name)

	return user
}

func FindUser(db *octmysql.OctMysql, id string) *User {

	row := db.QueryRow("SELECT ID,U_Name,U_State,U_Type,U_Email,U_PhoneNumber,"+
		"U_Description,U_CreateTime,U_LastLogin,U_LastSync "+
		"FROM tb_user WHERE ID = ? LIMIT 1", id)

	user := new(User)

	err := row.Scan(&user.Id, &user.Name, &user.State,
		&user.Type, &user.Email, &user.PhoneNumber, &user.Desc,
		&user.CreateTime, &user.LastLogin, &user.LastSync)
	if err != nil {
		logger.Errorf("Find account %s error %s", id, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", user.Id, user.Name)

	return user
}
