package command

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/turistikrota/service.account/src/domain/account"
	"github.com/turistikrota/service.account/src/domain/platform"
)

type AccountSocialAddCommand struct {
	UserUUID    string
	AccountName string
	Platform    string
	Value       string
}

type AccountSocialAddResult struct{}

type AccountSocialAddHandler decorator.CommandHandler[AccountSocialAddCommand, *AccountSocialAddResult]

type accountSocialAddHandler struct {
	platformRepo    platform.Repository
	platformFactory platform.Factory
	accountRepo     account.Repository
	accountFactory  account.Factory
	events          account.Events
}

type AccountSocialAddHandlerConfig struct {
	PlatformRepo    platform.Repository
	PlatformFactory platform.Factory
	AccountRepo     account.Repository
	AccountFactory  account.Factory
	Events          account.Events
	CqrsBase        decorator.Base
}

func NewAccountSocialAddHandler(config AccountSocialAddHandlerConfig) AccountSocialAddHandler {
	return decorator.ApplyCommandDecorators[AccountSocialAddCommand, *AccountSocialAddResult](
		accountSocialAddHandler{
			platformRepo:    config.PlatformRepo,
			platformFactory: config.PlatformFactory,
			accountRepo:     config.AccountRepo,
			accountFactory:  config.AccountFactory,
			events:          config.Events,
		},
		config.CqrsBase,
	)
}

func (h accountSocialAddHandler) Handle(ctx context.Context, command AccountSocialAddCommand) (*AccountSocialAddResult, *i18np.Error) {
	p, err := h.platformRepo.GetBySlug(ctx, command.Platform)
	if err != nil {
		return nil, err
	}
	err = h.platformFactory.ValidatePlatformValue(p, command.Value)
	if err != nil {
		return nil, err
	}
	social := &account.EntitySocial{
		Platform:   command.Platform,
		Value:      command.Value,
		FixedValue: h.platformFactory.FixPlatformValue(p, command.Value),
	}
	err = h.accountRepo.SocialAdd(ctx, account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	}, social)
	if err != nil {
		return nil, err
	}
	h.events.SocialAdded(account.UserUnique{
		UUID: command.UserUUID,
		Name: command.AccountName,
	}, *social)
	return nil, nil
}