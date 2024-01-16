package service

import (
	"github.com/cilloparch/cillop/events"
	"github.com/mixarchitecture/microp/validator"
	"github.com/turistikrota/service.account/app"
	"github.com/turistikrota/service.account/app/command"
	"github.com/turistikrota/service.account/app/query"
	"github.com/turistikrota/service.account/config"
	"github.com/turistikrota/service.account/domains/account"
	"github.com/turistikrota/service.shared/db/mongo"
)

type Config struct {
	App         config.App
	EventEngine events.Engine
	Mongo       *mongo.DB
	Validator   *validator.Validator
}

func NewApplication(config Config) app.Application {
	accountFactory := account.NewFactory()
	accountRepo := account.NewRepo(config.Mongo.GetCollection(config.App.DB.Account.Collection), accountFactory)
	accountEvents := account.NewEvents(account.EventConfig{
		Topics:    config.App.Topics,
		Publisher: config.EventEngine,
	})

	return app.Application{
		Commands: app.Commands{
			AccountDelete:  command.NewAccountDeleteHandler(accountRepo, accountEvents),
			AccountCreate:  command.NewAccountCreateHandler(accountFactory, accountRepo, accountEvents),
			AccountUpdate:  command.NewAccountUpdateHandler(accountFactory, accountRepo, accountEvents),
			AccountEnable:  command.NewAccountEnableHandler(accountRepo, accountEvents),
			AccountDisable: command.NewAccountDisableHandler(accountRepo, accountEvents),
		},
		Queries: app.Queries{
			AccountFilter:      query.NewAccountFilterHandler(accountRepo),
			AccountGet:         query.NewAccountGetHandler(accountRepo),
			AccountProfileView: query.NewAccountProfileViewHandler(accountRepo),
			AccountListMy:      query.NewAccountListMyHandler(accountRepo),
			AccountListByIds:   query.NewAccountListByIdsHandler(accountRepo),
		},
	}
}
