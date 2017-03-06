package account

import (
	"fmt"
	"octlink/mirage/src/utils/config"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
)

var logger *octlog.LogConfig

func InitLog(level int) {
	logger = octlog.InitLogConfig("account.log", level)
}

type Account struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	State       int    `json:"state"` // 1 for enable,0 for disabled
	Type        int    `json:"type"`  // 7 for superadmin, 3 for admin, 1 for audit
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"passord"`
	LastLogin   int64  `json:"lastLogin"`
	CreateTime  int64  `json:"createTime"`
	LastSync    int64  `json:"lastSync"`
	Desc        string `json:"desc"`
}

func (account *Account) Brief() map[string]interface{} {

	u := make(map[string]interface{}, 2)

	u["id"] = account.Id
	u["name"] = account.Name

	return u
}

func (account *Account) Add(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("INSERT INTO %s (ID, U_Name, U_Type, U_Email, U_PhoneNumber,"+
		"U_Password, U_CreateTime, U_Description) "+
		"VALUES ('%s', '%s', '%d')",
		config.TB_ACCOUNT, account.Id, account.Name, account.Type,
		account.Email, account.PhoneNumber, account.Password,
		account.CreateTime, account.Desc)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("insert db error %s", sql)
		return merrors.ERR_DB_ERR
	}

	return 0
}

func (account *Account) Delete(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("DELETE FROM %s WHERE ID='%s'", config.TB_ACCOUNT, account.Id)
	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("delete account %s error", account.Id)
		return merrors.ERR_DB_ERR
	}

	octlog.Debug(sql)

	return 0
}
