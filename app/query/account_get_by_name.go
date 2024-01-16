package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountGetByNameQuery struct {
	UserName string `params:"userName" validate:"required,username"`
}

type AccountGetByNameResult struct {
	Dto account.AdminDetailDto
}

type AccountGetByNameHandler cqrs.HandlerFunc[AccountGetByNameQuery, *AccountGetByNameResult]

func NewAccountGetByNameHandler(repo account.Repo) AccountGetByNameHandler {
	return func(ctx context.Context, query AccountGetByNameQuery) (*AccountGetByNameResult, *i18np.Error) {
		e, err := repo.GetByName(ctx, query.UserName)
		if err != nil {
			return nil, err
		}
		return &AccountGetByNameResult{
			Dto: e.ToAdminDetail(),
		}, nil
	}
}
