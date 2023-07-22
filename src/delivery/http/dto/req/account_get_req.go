package req

import "github.com/turistikrota/service.account/src/app/query"

type AccountGetRequest struct{}

func (r *AccountGetRequest) ToQuery(userUUID string, userName string) query.AccountGetQuery {
	return query.AccountGetQuery{
		UserUUID: userUUID,
		Name:     userName,
	}
}
