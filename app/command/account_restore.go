package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountRestoreCmd struct {
	UserName string `params:"userName" validate:"required,username"`
}

type AccountRestoreRes struct{}

type AccountRestoreHandler cqrs.HandlerFunc[AccountRestoreCmd, *AccountRestoreRes]

func NewAccountRestoreHandler(repo account.Repo, events account.Events) AccountRestoreHandler {
	return func(ctx context.Context, cmd AccountRestoreCmd) (*AccountRestoreRes, *i18np.Error) {
		res, err := repo.GetByName(ctx, cmd.UserName)
		if err != nil {
			return nil, err
		}
		err = repo.Restore(ctx, cmd.UserName)
		if err != nil {
			return nil, err
		}
		events.Restored(account.UserUnique{
			UUID: res.UserUUID,
			Name: cmd.UserName,
		})
		return &AccountRestoreRes{}, nil
	}
}
