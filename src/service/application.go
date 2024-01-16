package service

import (
	"github.com/mixarchitecture/cache"
	"github.com/mixarchitecture/microp/decorator"
	"github.com/mixarchitecture/microp/events"
	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.account/src/adapters"
	"github.com/turistikrota/service.account/src/app"
	"github.com/turistikrota/service.account/src/app/command"
	"github.com/turistikrota/service.account/src/app/query"
	"github.com/turistikrota/service.account/src/config"
	"github.com/turistikrota/service.account/src/domain/account"
	"github.com/turistikrota/service.account/src/domain/platform"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Mongo       *mongo.DB
	Validator   *validator.Validator
	CacheSrv    cache.Service
}

func NewApplication(config Config) app.Application {
	accountFactory := account.NewFactory()
	accountRepo := adapters.Mongo.NewAccount(accountFactory, config.Mongo.GetCollection(config.App.DB.Account.Collection))
	accountEvents := account.NewEvents(account.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
	})

	platformFactory := platform.NewFactory()
	platformRepo := adapters.Mongo.NewPlatform(platformFactory, config.Mongo.GetCollection(config.App.DB.Platform.Collection))
	platformEvents := platform.NewEvents(platform.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
	})

	base := decorator.NewBase()

	return app.Application{
		Commands: app.Commands{
			AccountCreate: command.NewAccountCreateHandler(command.AccountCreateHandlerConfig{
				Repo:     accountRepo,
				Factory:  accountFactory,
				Events:   accountEvents,
				CqrsBase: base,
			}),
			AccountUpdate: command.NewAccountUpdateHandler(command.AccountUpdateHandlerConfig{
				Repo:     accountRepo,
				Factory:  accountFactory,
				Events:   accountEvents,
				CqrsBase: base,
			}),
			AccountDelete: command.NewAccountDeleteHandler(command.AccountDeleteHandlerConfig{
				Repo:     accountRepo,
				Events:   accountEvents,
				Factory:  accountFactory,
				CqrsBase: base,
			}),
			AccountDisable: command.NewAccountDisableHandler(command.AccountDisableHandlerConfig{
				Repo:     accountRepo,
				Events:   accountEvents,
				Factory:  accountFactory,
				CqrsBase: base,
			}),
			AccountEnable: command.NewAccountEnableHandler(command.AccountEnableHandlerConfig{
				Repo:     accountRepo,
				Events:   accountEvents,
				Factory:  accountFactory,
				CqrsBase: base,
			}),
			AccountSocialAdd: command.NewAccountSocialAddHandler(command.AccountSocialAddHandlerConfig{
				PlatformRepo:    platformRepo,
				PlatformFactory: platformFactory,
				AccountRepo:     accountRepo,
				AccountFactory:  accountFactory,
				Events:          accountEvents,
				CqrsBase:        base,
			}),
			AccountSocialRemove: command.NewAccountSocialRemoveHandler(command.AccountSocialRemoveHandlerConfig{
				PlatformRepo:    platformRepo,
				PlatformFactory: platformFactory,
				AccountRepo:     accountRepo,
				AccountFactory:  accountFactory,
				Events:          accountEvents,
				CqrsBase:        base,
			}),
			AccountSocialUpdate: command.NewAccountSocialUpdateHandler(command.AccountSocialUpdateHandlerConfig{
				PlatformRepo:    platformRepo,
				PlatformFactory: platformFactory,
				AccountRepo:     accountRepo,
				AccountFactory:  accountFactory,
				Events:          accountEvents,
				CqrsBase:        base,
			}),
			PlatformCreate: command.NewPlatformCreateHandler(command.PlatformCreateHandlerConfig{
				Repo:     platformRepo,
				Factory:  platformFactory,
				Events:   platformEvents,
				CqrsBase: base,
			}),
			PlatformUpdate: command.NewPlatformUpdateHandler(command.PlatformUpdateHandlerConfig{
				Repo:     platformRepo,
				Factory:  platformFactory,
				Events:   platformEvents,
				CqrsBase: base,
			}),
			PlatformDelete: command.NewPlatformDeleteHandler(command.PlatformDeleteHandlerConfig{
				Repo:     platformRepo,
				Factory:  platformFactory,
				Events:   platformEvents,
				CqrsBase: base,
			}),
			PlatformDisable: command.NewPlatformDisableHandler(command.PlatformDisableHandlerConfig{
				Repo:     platformRepo,
				Factory:  platformFactory,
				Events:   platformEvents,
				CqrsBase: base,
			}),
			PlatformEnable: command.NewPlatformEnableHandler(command.PlatformEnableHandlerConfig{
				Repo:     platformRepo,
				Factory:  platformFactory,
				Events:   platformEvents,
				CqrsBase: base,
			}),
			PlatformTranslationCreate: command.NewPlatformTranslationCreateHandler(command.PlatformTranslationCreateHandlerConfig{
				Repo:     platformRepo,
				Factory:  platformFactory,
				Events:   platformEvents,
				CqrsBase: base,
			}),
			PlatformTranslationUpdate: command.NewPlatformTranslationUpdateHandler(command.PlatformTranslationUpdateHandlerConfig{
				Repo:     platformRepo,
				Factory:  platformFactory,
				Events:   platformEvents,
				CqrsBase: base,
			}),
			PlatformTranslationRemove: command.NewPlatformTranslationRemoveHandler(command.PlatformTranslationRemoveHandlerConfig{
				Repo:     platformRepo,
				Factory:  platformFactory,
				Events:   platformEvents,
				CqrsBase: base,
			}),
		},
		Queries: app.Queries{
			AccountGet: query.NewAccountGetHandler(query.AccountGetHandlerConfig{
				Repo:     accountRepo,
				CqrsBase: base,
				CacheSrv: config.CacheSrv,
			}),
			AccountListMy: query.NewAccountListMyHandler(query.AccountListMyHandlerConfig{
				Repo:     accountRepo,
				CqrsBase: base,
			}),
			AccountListByIds: query.NewAccountListByIdsHandler(query.AccountListByIdsHandlerConfig{
				Repo:     accountRepo,
				CqrsBase: base,
			}),
			AccountProfileView: query.NewAccountProfileViewHandler(query.AccountProfileViewHandlerConfig{
				Repo:     accountRepo,
				CqrsBase: base,
			}),
			PlatformGetBySlug: query.NewPlatformGetBySlugHandler(query.PlatformGetBySlugHandlerConfig{
				Repo:     platformRepo,
				CqrsBase: base,
			}),
			PlatformListAll: query.NewPlatformListAllHandler(query.PlatformListAllHandlerConfig{
				Repo:     platformRepo,
				CqrsBase: base,
			}),
		},
	}
}