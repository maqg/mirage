package session

import (
	"fmt"
	"octlink/mirage/src/utils/config"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
	"octlink/mirage/src/utils/uuid"
	"time"
)

const (
	SESSION_TIMEOUT = 2 * 60 * 60
)

var logger *octlog.LogConfig

func InitLog(level int) {
	logger = octlog.InitLogConfig("session.log", level)
}

type Session struct {
	Id         string `json:"id"`
	User       string `json:"user"`
	CreateTime int64  `json:"createTime"`
	LastSync   int64  `json:"lastSync"`
	ExpireTime int64  `json:"expireTime"`
	Cookie     string `json:"cookie"`
}

func (session *Session) Insert(db *octmysql.OctMysql) int {
	sql := fmt.Sprintf("INSERT INTO %s (ID,S_User,S_CreateTime,"+
		"S_LastSync,S_ExpireTime) VALUES ('%s','%s','%d','%d','%d');",
		config.TB_SESSION,
		session.Id,
		session.User,
		session.CreateTime,
		session.LastSync,
		session.ExpireTime)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("insert db error %s", sql)
		return merrors.ERR_DB_ERR
	}

	return 0
}

func GetSession(db *octmysql.OctMysql) *Session {
	sess := new(Session)

	now := int64(time.Now().Unix())

	sess.CreateTime = now
	sess.LastSync = now
	sess.ExpireTime = now + SESSION_TIMEOUT
	sess.Id = uuid.Generate().Simple()
	sess.User = "admin"

	sess.Insert(db)

	return sess
}
