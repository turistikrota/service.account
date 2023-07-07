package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.account/src/domain/account"
)

type AccountListMyQuery struct {
	UserUUID string
}

type AccountListMyResult struct {
	Entities []*account.Entity
}

type AccountListMyHandler decorator.QueryHandler[AccountListMyQuery, *AccountListMyResult]

type accountListMyHandler struct {
	repo account.Repository
}

type AccountListMyHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountListMyHandler(config AccountListMyHandlerConfig) AccountListMyHandler {
	return decorator.ApplyQueryDecorators[AccountListMyQuery, *AccountListMyResult](
		accountListMyHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h accountListMyHandler) Handle(ctx context.Context, query AccountListMyQuery) (*AccountListMyResult, *i18np.Error) {
	accounts, err := h.repo.ListMy(ctx, query.UserUUID)
	if err != nil {
		return nil, err
	}
	return &AccountListMyResult{
		Entities: accounts,
	}, nil
}
