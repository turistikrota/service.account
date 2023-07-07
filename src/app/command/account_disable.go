package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.account/src/domain/account"
)

type AccountDisableCommand struct {
	UserUUID    string
	AccountName string
}

type AccountDisableResult struct{}

type AccountDisableHandler decorator.CommandHandler[AccountDisableCommand, *AccountDisableResult]

type accountDisableHandler struct {
	repo    account.Repository
	factory account.Factory
	events  account.Events
}

type AccountDisableHandlerConfig struct {
	Repo     account.Repository
	Factory  account.Factory
	Events   account.Events
	CqrsBase decorator.Base
}

func NewAccountDisableHandler(config AccountDisableHandlerConfig) AccountDisableHandler {
	return decorator.ApplyCommandDecorators[AccountDisableCommand, *AccountDisableResult](
		accountDisableHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h accountDisableHandler) Handle(ctx context.Context, command AccountDisableCommand) (*AccountDisableResult, *i18np.Error) {
	err := h.repo.Disable(ctx, account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	})
	if err != nil {
		return nil, err
	}
	h.events.Disabled(account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	})
	return nil, nil
}
