package containers

import (
	"go-test/pkg/application/controllers"
	"go-test/pkg/domain/services"
	"go-test/pkg/interfaces"
	"go-test/pkg/persistence/repositories"
	"go-test/pkg/presentation/routers"
)

type ClientContainer struct {
	Router     routers.ClientRouter
	Service    services.ClientService
	Repository repositories.ClientRepository
}

func NewClientContainer(i interfaces.ClientInterface) ClientContainer {
	var repository = &repositories.ClientRepository{ClientInterface: i}
	var service = services.ClientService{Repository: *repository}
	var cc ClientContainer = ClientContainer{
		Repository: *repository,
		Service:    service,
		Router: routers.ClientRouter{
			Controller: controllers.ClientController{
				Service: service,
			},
		},
	}
	cc.Router.Router = cc.Router.NewClientRouter()
	return cc
}
