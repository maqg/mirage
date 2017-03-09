package api

import (
	"fmt"
	"octlink/mirage/src/modules/usergroup"
	"octlink/mirage/src/utils/merrors"
	"octlink/mirage/src/utils/octlog"
	"octlink/mirage/src/utils/uuid"
)

func APIAddUserGroup(paras *ApiParas) *ApiResponse {
	resp := new(ApiResponse)

	newGroup := usergroup.FindGroupByName(paras.Db, paras.InParas.Paras["name"].(string))
	if newGroup != nil {
		logger.Errorf("user %s already exist\n", newGroup.Name)
		resp.Error = merrors.ERR_SEGMENT_ALREADY_EXIST
		return resp
	}

	newGroup = new(usergroup.UserGroup)
	newGroup.Id = uuid.Generate().Simple()
	newGroup.Name = paras.InParas.Paras["name"].(string)
	newGroup.Desc = paras.InParas.Paras["desc"].(string)
	newGroup.AccountId = paras.InParas.Paras["accountId"].(string)

	resp.Error = newGroup.Add(paras.Db)

	return resp
}

func APIShowUserGroup(paras *ApiParas) *ApiResponse {

	octlog.Debug("running in APIShowUser\n")

	resp := new(ApiResponse)

	groupId := paras.InParas.Paras["id"].(string)
	temp := usergroup.FindGroup(paras.Db, groupId)

	if temp == nil {
		resp.Error = merrors.ERR_SEGMENT_NOT_EXIST
		resp.ErrorLog = fmt.Sprintf("group %s not found", groupId)
		return resp
	}

	resp.Data = temp

	octlog.Debug("found User %s", temp.Name)

	return resp
}

func APIUpdateUserGroup(paras *ApiParas) *ApiResponse {
	octlog.Debug("running in APIUpdateUser\n")
	resp := new(ApiResponse)
	resp.Error = 0
	return resp
}

func APIShowAllUserGroup(paras *ApiParas) *ApiResponse {

	resp := new(ApiResponse)

	offset := ParasInt(paras.InParas.Paras["start"])
	limit := ParasInt(paras.InParas.Paras["limit"])

	rows, err := paras.Db.Query("SELECT ID,UG_Name,UG_AccountId,"+
		"UG_CreateTime,UG_LastSync,UG_Description "+
		"FROM tb_usergroup LIMIT ?,?", offset, limit)
	if err != nil {
		logger.Errorf("query user group error %s\n", err.Error())
		resp.Error = merrors.ERR_DB_ERR
		return resp
	}
	defer rows.Close()

	groupList := make([]usergroup.UserGroup, 0)

	for rows.Next() {
		var group usergroup.UserGroup
		err = rows.Scan(&group.Id, &group.Name, &group.AccountId,
			&group.CreateTime, &group.LastSync, &group.Desc)
		if err == nil {
			logger.Debugf("query result: %s:%s\n", group.Id, group.Name)
		} else {
			logger.Errorf("query usergroup list error %s\n", err.Error())
		}
		groupList = append(groupList, group)
	}

	count := usergroup.GetGroupCount(paras.Db)

	result := make(map[string]interface{}, 3)
	result["total"] = count
	result["count"] = len(groupList)
	result["data"] = groupList

	resp.Data = result

	return resp
}

func APIDeleteUserGroup(paras *ApiParas) *ApiResponse {

	octlog.Debug("running in APIDeleteUser\n")

	resp := new(ApiResponse)

	group := usergroup.FindGroup(paras.Db, paras.InParas.Paras["id"].(string))
	if group == nil {
		resp.Error = merrors.ERR_SEGMENT_NOT_EXIST
		return resp
	}

	group.Delete(paras.Db)

	return resp
}

func APIShowUserGroupList(paras *ApiParas) *ApiResponse {

	resp := new(ApiResponse)

	rows, err := paras.Db.Query("SELECT ID,UG_Name FROM tb_usergroup")
	if err != nil {
		logger.Errorf("query user group list error %s\n", err.Error())
		resp.Error = merrors.ERR_DB_ERR
		return resp
	}
	defer rows.Close()

	groupList := make([]map[string]interface{}, 0)

	for rows.Next() {
		var group usergroup.UserGroup
		err = rows.Scan(&group.Id, &group.Name)
		if err == nil {
			logger.Debugf("query result: %s:%s\n", group.Id, group.Name)
		} else {
			logger.Errorf("query usergroup list error %s\n", err.Error())
		}
		groupList = append(groupList, group.Brief())
	}

	resp.Data = groupList

	return resp
}
