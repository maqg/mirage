package api

import (
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/octmysql"
	"testing"
)

func MakeParas(inParas *InputParas) *ApiParas {

	var apiParas *ApiParas = new(ApiParas)

	apiParas.InParas = inParas

	proto := FindApiProto(apiParas.InParas.Api)
	if proto == nil {
		octlog.Error("no api proto found for %s\n",
			apiParas.InParas.Api)
		return nil
	}

	apiParas.Proto = proto
	apiParas.User = "root"
	apiParas.Db = new(octmysql.OctMysql)

	return apiParas
}

func TestTest(t *testing.T) {
	t.Log("test OK")
}

func TestAdd(t *testing.T) {

	inParas := new(InputParas)

	inParas.Module = "user"
	inParas.Api = "octlink.mirage.center.user.APIAddAccount"

	// paras
	inParas.Paras = make(map[string]interface{}, 4)
	inParas.Paras["account"] = "test"
	inParas.Paras["password"] = "test"
	inParas.Paras["email"] = "test@test"
	inParas.Paras["phoneNumber"] = "12345678901"

	paras := MakeParas(inParas)

	resp := APIAddAccount(paras)

	if resp.Error != 0 {
		t.Error("Test Add Account Error")
	} else {
		t.Log("Add User OK")
	}
}

func TestDelete(t *testing.T) {
	inParas := new(InputParas)

	db := new(octmysql.OctMysql)
	user := FindUserByName(db, "test")
	if user == nil {
		t.Log("User test not exist")
		return
	}

	defer db.Close()

	inParas.Module = "user"
	inParas.Api = "octlink.mirage.center.user.APIDeleteAccount"

	// paras
	inParas.Paras = make(map[string]interface{}, 4)
	inParas.Paras["id"] = user.Id

	paras := MakeParas(inParas)

	resp := APIDeleteAccount(paras)

	if resp.Error != 0 {
		t.Error("Test Delte Account Error")
	} else {
		t.Log("Test Delete User OK")
	}
}

func init() {
	InitApiLog(octlog.DEBUG_LEVEL)
	octlog.InitDebugConfig(octlog.DEBUG_LEVEL)
	LoadApiConfig("../../")
}
