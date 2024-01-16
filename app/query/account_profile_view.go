package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountProfileViewQuery struct {
	UserName string `param:"userName" validate:"required,username"`
}

type AccountProfileViewResult struct {
	Dto account.ProfileDto
}

type AccountProfileViewHandler cqrs.HandlerFunc[AccountProfileViewQuery, *AccountProfileViewResult]

func NewAccountProfileViewHandler(repo account.Repo) AccountProfileViewHandler {
	return func(ctx context.Context, query AccountProfileViewQuery) (*AccountProfileViewResult, *i18np.Error) {
		e, err := repo.ProfileView(ctx, query.UserName)
		if err != nil {
			return nil, err
		}
		return &AccountProfileViewResult{
			Dto: e.ToProfile(),
		}, nil
	}
}
