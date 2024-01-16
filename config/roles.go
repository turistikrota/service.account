package config

import "github.com/turistikrota/service.shared/base_roles"

type roles struct {
	base_roles.Roles
	Account AccountRole
}

type AccountRole struct {
	List    string
	View    string
	Delete  string
	Restore string
	Super   string
}

var Roles = roles{
	Roles: base_roles.BaseRoles,
	Account: AccountRole{
		List:    "account.list",
		View:    "account.view",
		Delete:  "account.delete",
		Restore: "account.restore",
		Super:   "account.super",
	},
}
