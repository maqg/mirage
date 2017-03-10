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
	SESSION_TIMEOUT    = 2 * 60 * 60
	SESSION_DEFAULT_ID = "00000000000000000000000000000000"
)

var logger *octlog.LogConfig

func InitLog(level int) {
	logger = octlog.InitLogConfig("session.log", level)
}

type Session struct {
	Id         string `json:"id"`
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	CreateTime int64  `json:"createTime"`
	LastSync   int64  `json:"lastSync"`
	ExpireTime int64  `json:"expireTime"`
	Cookie     string `json:"cookie"`
}

func (session *Session) Insert(db *octmysql.OctMysql) int {

	sql := fmt.Sprintf("INSERT INTO %s (ID,S_UserId,S_UserName,S_CreateTime,"+
		"S_LastSync,S_ExpireTime) VALUES ('%s','%s','%s','%d','%d','%d');",
		config.TB_SESSION,
		session.Id,
		session.UserId,
		session.UserName,
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

func (session *Session) Delete(db *octmysql.OctMysql) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE ID='%s';",
		config.TB_SESSION, session.Id)
	db.Exec(sql)
}

func NewSession(db *octmysql.OctMysql, userId string, userName string) *Session {

	sess := new(Session)

	now := int64(time.Now().Unix())

	sess.Id = uuid.Generate().Simple()

	sess.CreateTime = now
	sess.LastSync = now
	sess.ExpireTime = now + SESSION_TIMEOUT
	sess.UserName = userName
	sess.UserId = userId

	err := sess.Insert(db)
	if err != 0 {
		octlog.Error("Insert new session %s,%s Error", userId, userName)
		return nil
	}

	return sess
}

func ClearSession(db *octmysql.OctMysql, userId string) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE S_UserId='%s';",
		config.TB_SESSION, userId)
	db.Exec(sql)
}

func (session *Session) Update(db *octmysql.OctMysql, sid string) int {

	if sid == SESSION_DEFAULT_ID {
		octlog.Debug("default session no need update %s", sid)
		return 0
	}

	sql := fmt.Sprintf("UPDATE %s SET S_LastSync='%d',S_ExpireTime='%d' "+
		"WHERE ID='%s';", config.TB_SESSION,
		time.Now().Unix(),
		time.Now().Unix()+SESSION_TIMEOUT,
		sid)

	_, err := db.Exec(sql)
	if err != nil {
		logger.Errorf("update session error %s", sql)
		return merrors.ERR_DB_ERR
	}

	octlog.Debug("session "+sid+" updated %s", sql)

	return 0
}

func FindSession(db *octmysql.OctMysql, sid string) *Session {

	sql := fmt.Sprintf("SELECT ID,S_UserId,S_UserName,S_Cookie,"+
		"S_CreateTime,S_LastSync,S_ExpireTime "+
		"FROM %s WHERE ID='%s' LIMIT 1;", config.TB_SESSION, sid)
	row := db.QueryRow(sql)

	session := new(Session)

	err := row.Scan(&session.Id, &session.UserId, &session.UserName,
		&session.Cookie, &session.CreateTime, &session.LastSync,
		&session.ExpireTime)
	if err != nil {
		logger.Errorf("Find session %s error %s, [%s]", sid,
			err.Error(), sql)
		return nil
	}

	session.Update(db, sid)

	octlog.Debug("found session, id:%s, user:%s", session.Id, session.UserName)

	return session
}
