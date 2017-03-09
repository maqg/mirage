package usergroup

import (
	"fmt"
	"octlink/mirage/src/utils/config"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
)

var logger *octlog.LogConfig

func InitLog(level int) {
	logger = octlog.InitLogConfig("usergroup.log", level)
}

type UserGroup struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	AccountId  string `json:"accountId"`
	CreateTime int64  `json:"createTime"`
	LastSync   int64  `json:"lastSync"`
	Desc       string `json:"desc"`
}

func GetGroupCount(db *octmysql.OctMysql) int {

	var count int = 0

	row := db.QueryRow("SELECT COUNT(1) FROM tb_usergroup")
	err := row.Scan(&count)
	if err != nil {
		logger.Errorf("get count for user group error %s", err.Error())
		return 0
	}

	octlog.Debug("got %d user groups", count)

	return count
}

func (group *UserGroup) Brief() map[string]interface{} {
	b := make(map[string]interface{}, 2)
	b["id"] = group.Id
	b["name"] = group.Name
	return b
}

func (group *UserGroup) Add(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("INSERT INTO %s (ID, UG_Name, UG_AccountId, "+
		"UG_CreateTime, UG_Description) VALUES ('%s', '%s', '%s', "+
		"'%d', '%s')",
		config.TB_USERGROUP,
		group.Id, group.Name, group.AccountId,
		group.CreateTime, group.Desc)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("insert db error %s", sql)
		return merrors.ERR_DB_ERR
	}

	return 0
}

func (group *UserGroup) Delete(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("DELETE FROM %s WHERE ID='%s'",
		config.TB_USERGROUP, group.Id)
	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("delete group %s error", group.Id)
		return merrors.ERR_DB_ERR
	}

	octlog.Debug(sql)

	return 0
}

func FindGroupByName(db *octmysql.OctMysql, name string) *UserGroup {

	row := db.QueryRow("SELECT ID,UG_Name,UG_AccountId,"+
		"UG_CreateTime,UG_LastSync, UG_Description "+
		"FROM tb_usergroup WHERE UG_Name = ? LIMIT 1", name)

	group := new(UserGroup)

	err := row.Scan(&group.Id, &group.Name, &group.AccountId,
		&group.CreateTime, &group.LastSync, &group.Desc)
	if err != nil {
		logger.Errorf("Find group group %s error %s", name, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", group.Id, group.Name)

	return group
}

func FindGroup(db *octmysql.OctMysql, id string) *UserGroup {

	row := db.QueryRow("SELECT ID,UG_Name,UG_AccountId,"+
		"UG_CreateTime,UG_LastSync, UG_Description "+
		"FROM tb_usergroup WHERE ID = ? LIMIT 1", id)

	group := new(UserGroup)

	err := row.Scan(&group.Id, &group.Name, &group.AccountId,
		&group.CreateTime, &group.LastSync, &group.Desc)
	if err != nil {
		logger.Errorf("Find group group %s error %s", id, err.Error())
		return nil
	}

	octlog.Debug("id %s, name :%s", group.Id, group.Name)

	return group
}
