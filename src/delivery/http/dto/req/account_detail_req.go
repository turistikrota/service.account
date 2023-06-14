package req

import (
	"api.turistikrota.com/account/src/app/command"
	"api.turistikrota.com/account/src/app/query"
	"github.com/turistikrota/service.shared/helper"
)

type AccountDetailRequest struct {
	UserNameAndCode string `param:"userNameAndCode" validate:"required,username_and_code"`
	UserName        string
	UserCode        string
}

func (r *AccountDetailRequest) Parse() *AccountDetailRequest {
	_, r.UserName, r.UserCode = helper.Parsers.ParseUsernameAndCode(r.UserNameAndCode)
	return r
}

func (r *AccountDetailRequest) ToDeleteCommand(userUUID string) command.AccountDeleteCommand {
	return command.AccountDeleteCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
		AccountCode: r.UserCode,
	}
}

func (r *AccountDetailRequest) ToDisableCommand(userUUID string) command.AccountDisableCommand {
	return command.AccountDisableCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
		AccountCode: r.UserCode,
	}
}

func (r *AccountDetailRequest) ToEnableCommand(userUUID string) command.AccountEnableCommand {
	return command.AccountEnableCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
		AccountCode: r.UserCode,
	}
}

func (r *AccountDetailRequest) ToGetQuery(userUUID string) query.AccountGetQuery {
	return query.AccountGetQuery{
		UserUUID: userUUID,
		Name:     r.UserName,
		Code:     r.UserCode,
	}
}

func (r *AccountDetailRequest) ToProfileQuery() query.AccountProfileViewQuery {
	return query.AccountProfileViewQuery{
		Name: r.UserName,
		Code: r.UserCode,
	}
}
