package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.account/src/domain/account"
)

type AccountDeleteCommand struct {
	UserUUID    string
	AccountName string
}

type AccountDeleteResult struct{}

type AccountDeleteHandler decorator.CommandHandler[AccountDeleteCommand, *AccountDeleteResult]

type accountDeleteHandler struct {
	repo    account.Repository
	factory account.Factory
	events  account.Events
}

type AccountDeleteHandlerConfig struct {
	Repo     account.Repository
	Factory  account.Factory
	Events   account.Events
	CqrsBase decorator.Base
}

func NewAccountDeleteHandler(config AccountDeleteHandlerConfig) AccountDeleteHandler {
	return decorator.ApplyCommandDecorators[AccountDeleteCommand, *AccountDeleteResult](
		accountDeleteHandler{
			repo:    config.Repo,
			factory: config.Factory,
			events:  config.Events,
		},
		config.CqrsBase,
	)
}

func (h accountDeleteHandler) Handle(ctx context.Context, command AccountDeleteCommand) (*AccountDeleteResult, *i18np.Error) {
	err := h.repo.Delete(ctx, account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	})
	if err != nil {
		return nil, err
	}
	h.events.Deleted(account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	})
	return nil, nil
}
