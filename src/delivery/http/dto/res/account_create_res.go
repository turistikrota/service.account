package res

import "github.com/turistikrota/service.account/src/app/command"

type AccountCreateResponse struct {
	AccountName string `json:"accountName"`
}

func (r *response) AccountCreate(res *command.AccountCreateResult) *AccountCreateResponse {
	return &AccountCreateResponse{
		AccountName: res.AccountName,
	}
}
