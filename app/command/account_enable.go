package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountEnableCmd struct {
	UserUUID string `params:"-"`
	UserName string `params:"-"`
}

type AccountEnableRes struct{}

type AccountEnableHandler cqrs.HandlerFunc[AccountEnableCmd, *AccountEnableRes]

func NewAccountEnableHandler(repo account.Repo, events account.Events) AccountEnableHandler {
	return func(ctx context.Context, cmd AccountEnableCmd) (*AccountEnableRes, *i18np.Error) {
		err := repo.Enable(ctx, account.UserUnique{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
		if err != nil {
			return nil, err
		}
		events.Enabled(account.UserUnique{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
		return &AccountEnableRes{}, nil
	}
}
