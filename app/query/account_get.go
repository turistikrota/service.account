package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountGetQuery struct {
	UserUUID string `params:"-"`
	UserName string `params:"-"`
}

type AccountGetResult struct {
	Dto account.DetailDto
}

type AccountGetHandler cqrs.HandlerFunc[AccountGetQuery, *AccountGetResult]

func NewAccountGetHandler(repo account.Repo) AccountGetHandler {
	return func(ctx context.Context, query AccountGetQuery) (*AccountGetResult, *i18np.Error) {
		e, err := repo.Get(ctx, account.UserUnique{
			UUID: query.UserUUID,
			Name: query.UserName,
		})
		if err != nil {
			return nil, err
		}
		return &AccountGetResult{
			Dto: e.ToDetail(),
		}, nil
	}
}
