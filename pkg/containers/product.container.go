package containers

import (
	"go-test/pkg/application/controllers"
	"go-test/pkg/domain/services"
	"go-test/pkg/interfaces"
	"go-test/pkg/persistence/repositories"
	"go-test/pkg/presentation/routers"
)

type ProductContainer struct {
	Router     routers.ProductRouter
	Service    services.ProductService
	Repository repositories.ProductRepository
}

func NewProductContainer(i interfaces.ProductInterface) ProductContainer {
	var repository = &repositories.ProductRepository{ProductInterface: i}
	var service = services.ProductService{Repository: *repository}
	var pc ProductContainer = ProductContainer{
		Repository: *repository,
		Service:    service,
		Router: routers.ProductRouter{
			Controller: controllers.ProductController{
				Service: service,
			},
		},
	}
	pc.Router.Router = pc.Router.NewProductRouter()
	return pc
}
