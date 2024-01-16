package rpc

import (
	"github.com/cilloparch/cillop/helpers/rpc"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/server"
	"github.com/turistikrota/service.account/app"
	protos "github.com/turistikrota/service.account/protos/account"
	"google.golang.org/grpc"
)

type srv struct {
	port int
	app  app.Application
	i18n i18np.I18n
	protos.AccountServiceServer
}

type Config struct {
	Port int
	App  app.Application
	I18n i18np.I18n
}

func New(cnf Config) server.Server {
	return srv{
		app:  cnf.App,
		i18n: cnf.I18n,
		port: cnf.Port,
	}
}

func (s srv) Listen() error {
	return rpc.RunServer(s.port, func(server *grpc.Server) {
		protos.RegisterAccountServiceServer(server, s)
	})
}
