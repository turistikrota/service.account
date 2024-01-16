package query

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	"github.com/turistikrota/service.account/domains/account"
	"github.com/turistikrota/service.account/pkg/utils"
)

type AccountFilterQuery struct {
	*utils.Pagination
	*account.FilterEntity
	UserUUID string `params:"-"`
	UserName string `params:"-"`
}

type AccountFilterResult struct {
	List *list.Result[account.AdminListDto]
}

type AccountFilterHandler cqrs.HandlerFunc[AccountFilterQuery, *AccountFilterResult]

func NewAccountFilterHandler(repo account.Repo) AccountFilterHandler {
	return func(ctx context.Context, query AccountFilterQuery) (*AccountFilterResult, *i18np.Error) {
		query.Default()
		offset := (*query.Page - 1) * *query.Limit
		accounts, err := repo.Filter(ctx, *query.FilterEntity, list.Config{
			Offset: offset,
			Limit:  *query.Limit,
		})
		if err != nil {
			return nil, err
		}
		res := make([]account.AdminListDto, len(accounts.List))
		for i, e := range accounts.List {
			res[i] = e.ToAdminList()
		}
		return &AccountFilterResult{
			List: &list.Result[account.AdminListDto]{
				Total:         accounts.Total,
				List:          res,
				FilteredTotal: accounts.FilteredTotal,
				Page:          accounts.Page,
				IsNext:        accounts.IsNext,
				IsPrev:        accounts.IsPrev,
			},
		}, nil
	}
}
