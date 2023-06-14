package command

import (
	"context"

	"api.turistikrota.com/account/src/domain/account"
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.shared/decorator"
)

type AccountDisableCommand struct {
	UserUUID    string
	AccountName string
	AccountCode string
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
		Code: command.AccountCode,
	})
	if err != nil {
		return nil, err
	}
	h.events.Disabled(account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
		Code: command.AccountCode,
	})
	return nil, nil
}
