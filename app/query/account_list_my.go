package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountListMyQuery struct {
	UserUUID string `params:"-"`
}

type AccountListMyResult struct {
	Dtos []account.ListDto
}

type AccountListMyHandler cqrs.HandlerFunc[AccountListMyQuery, *AccountListMyResult]

func NewAccountListMyHandler(repo account.Repo) AccountListMyHandler {
	return func(ctx context.Context, query AccountListMyQuery) (*AccountListMyResult, *i18np.Error) {
		accounts, err := repo.ListMy(ctx, query.UserUUID)
		if err != nil {
			return nil, err
		}
		res := make([]account.ListDto, len(accounts))
		for i, e := range accounts {
			res[i] = e.ToList()
		}
		return &AccountListMyResult{
			Dtos: res,
		}, nil
	}
}
