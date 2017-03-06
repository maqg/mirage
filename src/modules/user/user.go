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
