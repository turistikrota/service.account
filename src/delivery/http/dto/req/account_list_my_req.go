package req

import "github.com/turistikrota/service.account/src/app/query"

type AccountListMyRequest struct{}

func (r *AccountListMyRequest) ToQuery(userUUID string) query.AccountListMyQuery {
	return query.AccountListMyQuery{
		UserUUID: userUUID,
	}
}
