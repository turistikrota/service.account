package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountListAsClaimQuery struct {
	UserUUID string
}

type AccountListAsClaimResult struct {
	Entities []*account.Entity
}

type AccountListAsClaimHandler cqrs.HandlerFunc[AccountListAsClaimQuery, *AccountListAsClaimResult]

func NewAccountListAsClaimHandler(repo account.Repo) AccountListAsClaimHandler {
	return func(ctx context.Context, query AccountListAsClaimQuery) (*AccountListAsClaimResult, *i18np.Error) {
		e, err := repo.ListAsClaim(ctx, query.UserUUID)
		if err != nil {
			return nil, err
		}
		return &AccountListAsClaimResult{
			Entities: e,
		}, nil
	}
}
