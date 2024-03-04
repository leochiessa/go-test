package containers

import (
	"go-test/pkg/application/controllers"
	"go-test/pkg/domain/services"
	"go-test/pkg/interfaces"
	"go-test/pkg/persistence/repositories"
	"go-test/pkg/presentation/routers"
)

type RequestContainer struct {
	Router     routers.RequestRouter
	Service    services.RequestService
	Repository repositories.RequestRepository
}

func NewRequestContainer(i interfaces.RequestInterface) RequestContainer {
	var repository = &repositories.RequestRepository{RequestInterface: i}
	var service = services.RequestService{Repository: *repository}
	var rc RequestContainer = RequestContainer{
		Repository: *repository,
		Service:    service,
		Router: routers.RequestRouter{
			Controller: controllers.RequestController{
				Service: service,
			},
		},
	}
	rc.Router.Router = rc.Router.NewRequestRouter()
	return rc
}
