package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.account/domains/account"
)

type AccountCreateCmd struct {
	UserUUID string `json:"-"`
	UserName string `json:"userName" validate:"required,username,max=20,min=3"`
}

type AccountCreateRes struct{}

type AccountCreateHandler cqrs.HandlerFunc[AccountCreateCmd, *AccountCreateRes]

func NewAccountCreateHandler(factory account.Factory, repo account.Repo, events account.Events) AccountCreateHandler {
	return func(ctx context.Context, cmd AccountCreateCmd) (*AccountCreateRes, *i18np.Error) {
		u := account.UserUnique{
			UUID: cmd.UserUUID,
			Name: cmd.UserName,
		}
		exist, err := repo.Exist(ctx, u)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, factory.Errors.ErrAlreadyExist()
		}
		acc := factory.New(cmd.UserUUID, cmd.UserName)
		err = factory.Validate(acc)
		if err != nil {
			return nil, err
		}
		_, err = repo.Create(ctx, acc)
		if err != nil {
			return nil, err
		}
		events.Created(u)
		return &AccountCreateRes{}, nil
	}
}
