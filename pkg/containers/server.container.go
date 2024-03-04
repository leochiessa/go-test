package containers

import (
	"go-test/pkg/application/server"
	"go-test/pkg/presentation/routers"
	"os"
)

type ServerContainer struct {
	Server *server.Server
}

func NewServerContainer(cr routers.ClientRouter, pr routers.ProductRouter, rr routers.RequestRouter) (s ServerContainer) {
	var httpConnector *routers.AppRouter = &routers.AppRouter{
		ClientRouter:  cr,
		ProductRouter: pr,
		RequestRouter: rr,
	}
	httpConnector.Router = httpConnector.NewAppRouter()
	var sc ServerContainer = ServerContainer{&server.Server{
		AppRouter: httpConnector,
		HTTP_PORT: os.Getenv("HTTP_PORT"),
	}}
	return sc
}
