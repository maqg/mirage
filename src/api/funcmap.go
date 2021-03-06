// Auto generated by System, don't modify this file
package api

func InitApiService() {

	var service *ApiService

	GApiServices = make(map[string]*ApiService, 10000)

	// --------------------
	// for APIAddUser
	// --------------------
	service = new(ApiService)
	service.Name = "APIAddUser"
	service.Handler = APIAddUser
	GApiServices["octlink.mirage.center.user.APIAddUser"] = service

	// --------------------
	// for APIDeleteUser
	// --------------------
	service = new(ApiService)
	service.Name = "APIDeleteUser"
	service.Handler = APIDeleteUser
	GApiServices["octlink.mirage.center.user.APIDeleteUser"] = service

	// --------------------
	// for APILoginByUser
	// --------------------
	service = new(ApiService)
	service.Name = "APILoginByUser"
	service.Handler = APILoginByUser
	GApiServices["octlink.mirage.center.user.APILoginByUser"] = service

	// --------------------
	// for APIResetUserPassword
	// --------------------
	service = new(ApiService)
	service.Name = "APIResetUserPassword"
	service.Handler = APIResetUserPassword
	GApiServices["octlink.mirage.center.user.APIResetUserPassword"] = service

	// --------------------
	// for APIShowAllUser
	// --------------------
	service = new(ApiService)
	service.Name = "APIShowAllUser"
	service.Handler = APIShowAllUser
	GApiServices["octlink.mirage.center.user.APIShowAllUser"] = service

	// --------------------
	// for APIShowUser
	// --------------------
	service = new(ApiService)
	service.Name = "APIShowUser"
	service.Handler = APIShowUser
	GApiServices["octlink.mirage.center.user.APIShowUser"] = service

	// --------------------
	// for APIShowUserList
	// --------------------
	service = new(ApiService)
	service.Name = "APIShowUserList"
	service.Handler = APIShowUserList
	GApiServices["octlink.mirage.center.user.APIShowUserList"] = service

	// --------------------
	// for APIUpdateUser
	// --------------------
	service = new(ApiService)
	service.Name = "APIUpdateUser"
	service.Handler = APIUpdateUser
	GApiServices["octlink.mirage.center.user.APIUpdateUser"] = service

	// --------------------
	// for APIUpdateUserPassword
	// --------------------
	service = new(ApiService)
	service.Name = "APIUpdateUserPassword"
	service.Handler = APIUpdateUserPassword
	GApiServices["octlink.mirage.center.user.APIUpdateUserPassword"] = service

	// --------------------
	// for APIUserLogOut
	// --------------------
	service = new(ApiService)
	service.Name = "APIUserLogOut"
	service.Handler = APIUserLogOut
	GApiServices["octlink.mirage.center.user.APIUserLogOut"] = service

	// --------------------
	// for APIAddHost
	// --------------------
	service = new(ApiService)
	service.Name = "APIAddHost"
	service.Handler = APIAddHost
	GApiServices["octlink.mirage.center.host.APIAddHost"] = service

	// --------------------
	// for APIDeleteHost
	// --------------------
	service = new(ApiService)
	service.Name = "APIDeleteHost"
	service.Handler = APIDeleteHost
	GApiServices["octlink.mirage.center.host.APIDeleteHost"] = service

	// --------------------
	// for APIAccountLogOut
	// --------------------
	service = new(ApiService)
	service.Name = "APIAccountLogOut"
	service.Handler = APIAccountLogOut
	GApiServices["octlink.mirage.center.account.APIAccountLogOut"] = service

	// --------------------
	// for APIAddAccount
	// --------------------
	service = new(ApiService)
	service.Name = "APIAddAccount"
	service.Handler = APIAddAccount
	GApiServices["octlink.mirage.center.account.APIAddAccount"] = service

	// --------------------
	// for APIDeleteAccount
	// --------------------
	service = new(ApiService)
	service.Name = "APIDeleteAccount"
	service.Handler = APIDeleteAccount
	GApiServices["octlink.mirage.center.account.APIDeleteAccount"] = service

	// --------------------
	// for APILoginByAccount
	// --------------------
	service = new(ApiService)
	service.Name = "APILoginByAccount"
	service.Handler = APILoginByAccount
	GApiServices["octlink.mirage.center.account.APILoginByAccount"] = service

	// --------------------
	// for APIResetAccountPassword
	// --------------------
	service = new(ApiService)
	service.Name = "APIResetAccountPassword"
	service.Handler = APIResetAccountPassword
	GApiServices["octlink.mirage.center.account.APIResetAccountPassword"] = service

	// --------------------
	// for APIShowAccount
	// --------------------
	service = new(ApiService)
	service.Name = "APIShowAccount"
	service.Handler = APIShowAccount
	GApiServices["octlink.mirage.center.account.APIShowAccount"] = service

	// --------------------
	// for APIShowAccountList
	// --------------------
	service = new(ApiService)
	service.Name = "APIShowAccountList"
	service.Handler = APIShowAccountList
	GApiServices["octlink.mirage.center.account.APIShowAccountList"] = service

	// --------------------
	// for APIShowAllAccount
	// --------------------
	service = new(ApiService)
	service.Name = "APIShowAllAccount"
	service.Handler = APIShowAllAccount
	GApiServices["octlink.mirage.center.account.APIShowAllAccount"] = service

	// --------------------
	// for APIUpdateAccount
	// --------------------
	service = new(ApiService)
	service.Name = "APIUpdateAccount"
	service.Handler = APIUpdateAccount
	GApiServices["octlink.mirage.center.account.APIUpdateAccount"] = service

	// --------------------
	// for APIUpdateAccountPassword
	// --------------------
	service = new(ApiService)
	service.Name = "APIUpdateAccountPassword"
	service.Handler = APIUpdateAccountPassword
	GApiServices["octlink.mirage.center.account.APIUpdateAccountPassword"] = service

	// --------------------
	// for APIAddUserGroup
	// --------------------
	service = new(ApiService)
	service.Name = "APIAddUserGroup"
	service.Handler = APIAddUserGroup
	GApiServices["octlink.mirage.center.usergroup.APIAddUserGroup"] = service

	// --------------------
	// for APIDeleteUserGroup
	// --------------------
	service = new(ApiService)
	service.Name = "APIDeleteUserGroup"
	service.Handler = APIDeleteUserGroup
	GApiServices["octlink.mirage.center.usergroup.APIDeleteUserGroup"] = service

	// --------------------
	// for APIShowAllUserGroup
	// --------------------
	service = new(ApiService)
	service.Name = "APIShowAllUserGroup"
	service.Handler = APIShowAllUserGroup
	GApiServices["octlink.mirage.center.usergroup.APIShowAllUserGroup"] = service

	// --------------------
	// for APIShowUserGroup
	// --------------------
	service = new(ApiService)
	service.Name = "APIShowUserGroup"
	service.Handler = APIShowUserGroup
	GApiServices["octlink.mirage.center.usergroup.APIShowUserGroup"] = service

	// --------------------
	// for APIShowUserGroupList
	// --------------------
	service = new(ApiService)
	service.Name = "APIShowUserGroupList"
	service.Handler = APIShowUserGroupList
	GApiServices["octlink.mirage.center.usergroup.APIShowUserGroupList"] = service

	// --------------------
	// for APIUpdateUserGroup
	// --------------------
	service = new(ApiService)
	service.Name = "APIUpdateUserGroup"
	service.Handler = APIUpdateUserGroup
	GApiServices["octlink.mirage.center.usergroup.APIUpdateUserGroup"] = service
}
