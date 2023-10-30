package query

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.account/src/domain/account"
)

type AccountListByIdsQuery struct {
	Users []account.UserUnique
}

type AccountListByIdsResult struct {
	Entities []*account.Entity
}

type AccountListByIdsHandler decorator.QueryHandler[AccountListByIdsQuery, *AccountListByIdsResult]

type accountListByIdsHandler struct {
	repo account.Repository
}

type AccountListByIdsHandlerConfig struct {
	Repo     account.Repository
	CqrsBase decorator.Base
}

func NewAccountListByIdsHandler(config AccountListByIdsHandlerConfig) AccountListByIdsHandler {
	return decorator.ApplyQueryDecorators[AccountListByIdsQuery, *AccountListByIdsResult](
		accountListByIdsHandler{
			repo: config.Repo,
		},
		config.CqrsBase,
	)
}

func (h accountListByIdsHandler) Handle(ctx context.Context, query AccountListByIdsQuery) (*AccountListByIdsResult, *i18np.Error) {
	accounts, err := h.repo.ListByUniques(ctx, query.Users)
	if err != nil {
		return nil, err
	}
	return &AccountListByIdsResult{
		Entities: accounts,
	}, nil
}
