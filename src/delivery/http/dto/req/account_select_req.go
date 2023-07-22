package req

import "github.com/turistikrota/service.account/src/app/query"

type AccountSelectRequest struct {
	UserName string `param:"userName" validate:"required,username"`
}

func (r *AccountSelectRequest) ToGetQuery(userUUID string) query.AccountGetQuery {
	return query.AccountGetQuery{
		UserUUID: userUUID,
		Name:     r.UserName,
	}
}
