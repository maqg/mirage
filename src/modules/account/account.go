package account

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"octlink/mirage/src/modules/session"
	"octlink/mirage/src/utils/config"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
	"time"
)

var logger *octlog.LogConfig

func InitLog(level int) {
	logger = octlog.InitLogConfig("account.log", level)
}

func GetEncPassword(clearText string) string {

	m := md5.New()

	m.Write([]byte("Octopus"))
	m.Write([]byte(clearText))
	m.Write([]byte("Link"))

	enc := hex.EncodeToString(m.Sum(nil))

	octlog.Debug("got new enc password %s", enc)

	return enc
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

func GetAccountCount(db *octmysql.OctMysql) int {

	var count int = 0

	row := db.QueryRow("SELECT COUNT(1) FROM tb_account")
	err := row.Scan(&count)
	if err != nil {
		logger.Errorf("get count for account error %s", err.Error())
		return 0
	}

	octlog.Debug("got %d accounts", count)

	return count
}

func (account *Account) Brief() map[string]interface{} {
	b := make(map[string]interface{}, 2)
	b["id"] = account.Id
	b["name"] = account.Name
	return b
}

func (account *Account) ResetPassword(db *octmysql.OctMysql, password string) int {
	encPass := GetEncPassword(password)

	sql := fmt.Sprintf("UPDATE %s SET U_Password='%s',U_LastSync='%d' WHERE ID='%s';",
		config.TB_ACCOUNT, encPass, int64(time.Now().Unix()), account.Id)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("reset password for %s error %s",
			account.Name, sql)
		return merrors.ERR_DB_ERR
	}

	return 0
}

func (account *Account) Update(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("UPDATE %s SET U_Email='%s',U_PhoneNumber='%s', "+
		"U_Description='%s',U_LastSync='%d' WHERE ID='%s';",
		config.TB_ACCOUNT, account.Email,
		account.PhoneNumber, account.Desc,
		int64(time.Now().Unix()),
		account.Id)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("update account %s error %s",
			account.Name, sql)
		return merrors.ERR_DB_ERR
	}

	return 0
}

func (account *Account) UpdateLogin(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("UPDATE %s SET U_LastLogin='%d' WHERE ID='%s';",
		config.TB_ACCOUNT, int64(time.Now().Unix()), account.Id)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("update Login %s error %s",
			account.Name, sql)
		return merrors.ERR_DB_ERR
	}

	return 0
}

func (account *Account) UpdateSyncTime(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("UPDATE %s SET U_LastSync='%d' WHERE ID='%s';",
		config.TB_ACCOUNT, int64(time.Now().Unix()), account.Id)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("update last sync %s error %s",
			account.Name, sql)
		return merrors.ERR_DB_ERR
	}

	return 0
}

func (account *Account) Login(db *octmysql.OctMysql,
	password string) *session.Session {

	var accountId string

	encPass := GetEncPassword(password)

	sql := fmt.Sprintf("SELECT ID FROM %s WHERE U_Name='%s' AND U_Password='%s';",
		config.TB_ACCOUNT, account.Name, encPass)

	row := db.QueryRow(sql)
	err := row.Scan(&accountId)
	if err != nil {
		logger.Errorf("Login account %s error %s", account.Name, err.Error())
		return nil
	}

	account.UpdateLogin(db)

	return session.NewSession(db, account.Id, account.Name)
}

func (account *Account) Add(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("INSERT INTO %s (ID, U_Name, U_Type, "+
		"U_Email, U_PhoneNumber, U_Password, U_CreateTime, "+
		"U_Description) VALUES ('%s', '%s', '%d', '%s', '%s', "+
		"'%s', '%d', '%s')",
		config.TB_ACCOUNT,
		account.Id, account.Name, account.Type,
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

	sql := fmt.Sprintf("DELETE FROM %s WHERE ID='%s'",
		config.TB_ACCOUNT, account.Id)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("delete account %s error", account.Id)
		return merrors.ERR_DB_ERR
	}

	octlog.Debug(sql)

	return 0
}

func FindAccountByName(db *octmysql.OctMysql, name string) *Account {

	row := db.QueryRow("SELECT ID,U_Name,U_State,U_Type,U_Email,U_PhoneNumber,"+
		"U_Description,U_CreateTime,U_LastLogin,U_LastSync "+
		"FROM tb_account WHERE U_Name = ? LIMIT 1", name)

	account := new(Account)

	err := row.Scan(&account.Id, &account.Name, &account.State,
		&account.Type, &account.Email, &account.PhoneNumber, &account.Desc,
		&account.CreateTime, &account.LastLogin, &account.LastSync)
	if err != nil {
		logger.Errorf("Find account %s error %s", name, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", account.Id, account.Name)

	return account
}

func FindAccount(db *octmysql.OctMysql, id string) *Account {

	row := db.QueryRow("SELECT ID,U_Name,U_State,U_Type,U_Email,U_PhoneNumber,"+
		"U_Description,U_CreateTime,U_LastLogin,U_LastSync "+
		"FROM tb_account WHERE ID = ? LIMIT 1", id)

	account := new(Account)
	err := row.Scan(&account.Id, &account.Name, &account.State,
		&account.Type, &account.Email, &account.PhoneNumber, &account.Desc,
		&account.CreateTime, &account.LastLogin, &account.LastSync)
	if err != nil {
		logger.Errorf("Find account %s error %s", id, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", account.Id, account.Name)

	return account
}
