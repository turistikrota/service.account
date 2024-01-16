package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountDeleteCmd struct {
	UserUUID string `params:"-"`
	UserName string `params:"userName" validate:"required,username"`
}

type AccountDeleteRes struct{}

type AccountDeleteHandler cqrs.HandlerFunc[AccountDeleteCmd, *AccountDeleteRes]

func NewAccountDeleteHandler(repo account.Repo, events account.Events) AccountDeleteHandler {
	return func(ctx context.Context, cmd AccountDeleteCmd) (*AccountDeleteRes, *i18np.Error) {
		err := repo.Delete(ctx, account.UserUnique{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
		if err != nil {
			return nil, err
		}
		events.Deleted(account.UserUnique{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		})
		return &AccountDeleteRes{}, nil
	}
}
