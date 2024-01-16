package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountListByUserQuery struct {
	UserUUID string `params:"uuid" validate:"required,object_id"`
}

type AccountListByUserResult struct {
	Dtos []account.AdminListDto
}

type AccountListByUserHandler cqrs.HandlerFunc[AccountListByUserQuery, *AccountListByUserResult]

func NewAccountListByUserHandler(repo account.Repo) AccountListByUserHandler {
	return func(ctx context.Context, query AccountListByUserQuery) (*AccountListByUserResult, *i18np.Error) {
		e, err := repo.ListByUserId(ctx, query.UserUUID)
		if err != nil {
			return nil, err
		}
		res := make([]account.AdminListDto, len(e))
		for i, v := range e {
			res[i] = v.ToAdminList()
		}
		return &AccountListByUserResult{
			Dtos: res,
		}, nil
	}
}
