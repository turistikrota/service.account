package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountListByIdsQuery struct {
	Users []account.UserUnique
}

type AccountListByIdsResult struct {
	Entities []*account.Entity
}

type AccountListByIdsHandler cqrs.HandlerFunc[AccountListByIdsQuery, *AccountListByIdsResult]

func NewAccountListByIdsHandler(repo account.Repo) AccountListByIdsHandler {
	return func(ctx context.Context, query AccountListByIdsQuery) (*AccountListByIdsResult, *i18np.Error) {
		res, err := repo.ListByUniques(ctx, query.Users)
		if err != nil {
			return nil, err
		}
		return &AccountListByIdsResult{
			Entities: res,
		}, nil
	}
}
