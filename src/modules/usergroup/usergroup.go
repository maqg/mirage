package usergroup

import (
	"fmt"
	"octlink/mirage/src/utils/config"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
)

const (
	DEFAULT_USERGROUP = "00000000000000000000000000000000"
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
	count, _ := db.Count(config.TB_USERGROUP, "")
	return count
}

func (group *UserGroup) Brief() map[string]string {
	return map[string]string{
		"id":   group.Id,
		"name": group.Name,
	}
}

func (group *UserGroup) UserCount(db *octmysql.OctMysql) int {
	count, _ := db.Count(config.TB_RELUSERGROUP, "WHERE RUG_GroupId=?",
		group.Id)
	return count
}

func (group *UserGroup) Update(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("UPDATE %s SET UG_Name='%s',UG_Descrition='%s' "+
		"WHERE ID='%s';", config.TB_USERGROUP, group.Name,
		group.Desc, group.Id)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("Update UserGroup error %s", sql)
		return merrors.ERR_DB_ERR
	}

	return 0
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

	if group.Id == DEFAULT_USERGROUP {
		octlog.Error("default usergroup cannot be removed %s", DEFAULT_USERGROUP)
		return merrors.ERR_UNACCP_PARAS
	}

	if group.UserCount(db) != 0 {
		logger.Errorf("Before delete group, users should be null")
		return merrors.ERR_USERGROUP_USERS_NOT_EMPTY
	}

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
