package session

import (
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
	"testing"
)

func TestGetSession(t *testing.T) {

	db := new(octmysql.OctMysql)

	sess := GetSession(db)

	t.Log("session %s", sess.Id)

	t.Log("Test Get Session OK")
}

func init() {
	InitLog(octlog.DEBUG_LEVEL)
	octlog.InitDebugConfig(octlog.DEBUG_LEVEL)
}
