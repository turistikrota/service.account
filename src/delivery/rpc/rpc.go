package rpc

import (
	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.account/protos/account"
	"github.com/turistikrota/service.account/src/app"
)

type Server struct {
	app  app.Application
	i18n i18np.I18n
	account.AccountServiceServer
}

func New(a app.Application, i18n i18np.I18n) account.AccountServiceServer {
	return Server{
		app:  a,
		i18n: i18n,
	}
}
