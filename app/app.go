package app

import (
	"github.com/turistikrota/service.account/app/command"
	"github.com/turistikrota/service.account/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AccountDelete  command.AccountDeleteHandler
	AccountCreate  command.AccountCreateHandler
	AccountUpdate  command.AccountUpdateHandler
	AccountEnable  command.AccountEnableHandler
	AccountDisable command.AccountDisableHandler
}

type Queries struct {
	AccountFilter      query.AccountFilterHandler
	AccountGet         query.AccountGetHandler
	AccountProfileView query.AccountProfileViewHandler
	AccountListMy      query.AccountListMyHandler
	AccountListByIds   query.AccountListByIdsHandler
}
