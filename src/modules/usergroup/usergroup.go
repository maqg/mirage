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
	Id      string `json:"id"`
	Name    string `json:"name"`
	Account string `json:"accountId"`

	LastLogin  int64 `json:"lastLogin"`
	CreateTime int64 `json:"createTime"`
	LastSync   int64 `json:"lastSync"`
}

func (group *UserGroup) Brief() map[string]interface{} {
	b := make(map[string]interface{}, 2)
	b["id"] = group.Id
	b["name"] = group.Name
	return b
}

func (group *UserGroup) Add(db *octmysql.OctMysql) int {
	return 0
}

func (group *UserGroup) Delete(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("DELETE FROM %s WHERE ID='%s'", config.TB_USERGROUP, group.Id)
	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("delete group %s error", group.Id)
		return merrors.ERR_DB_ERR
	}

	octlog.Debug(sql)

	return 0
}
