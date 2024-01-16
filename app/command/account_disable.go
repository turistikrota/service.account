package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountDisableCmd struct {
	UserUUID string `params:"-"`
	UserName string `params:"userName" validate:"required,username"`
}

type AccountDisableRes struct{}

type AccountDisableHandler cqrs.HandlerFunc[AccountDisableCmd, *AccountDisableRes]

func NewAccountDisableHandler(repo account.Repo, events account.Events) AccountDisableHandler {
	return func(ctx context.Context, cmd AccountDisableCmd) (*AccountDisableRes, *i18np.Error) {
		err := repo.Disable(ctx, account.UserUnique{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
		if err != nil {
			return nil, err
		}
		events.Disabled(account.UserUnique{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
		return &AccountDisableRes{}, nil
	}
}
