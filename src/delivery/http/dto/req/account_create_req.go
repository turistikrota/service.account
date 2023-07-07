package req

import "github.com/turistikrota/service.account/src/app/command"

type AccountCreateRequest struct {
	UserName string `json:"userName" validate:"required,username,max=20,min=3"`
}

func (r *AccountCreateRequest) ToCommand(userUUID string) command.AccountCreateCommand {
	return command.AccountCreateCommand{
		UserUUID:    userUUID,
		AccountName: r.UserName,
	}
}
