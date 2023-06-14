package res

import "api.turistikrota.com/account/src/app/command"

type AccountCreateResponse struct {
	AccountName string `json:"accountName"`
	AccountCode string `json:"accountCode"`
}

func (r *response) AccountCreate(res *command.AccountCreateResult) *AccountCreateResponse {
	return &AccountCreateResponse{
		AccountName: res.AccountName,
		AccountCode: res.AccountCode,
	}
}
