package containers

import (
	"go-test/pkg/interfaces"
	"go-test/pkg/persistence/adapters"
	"go-test/pkg/persistence/config"
)

type DataBaseContainer struct {
	ClientInterface  interfaces.ClientInterface
	ProductInterface interfaces.ProductInterface
	RequestInterface interfaces.RequestInterface
}

func NewDataBaseContainer() *DataBaseContainer {
	var dbc *DataBaseContainer = &DataBaseContainer{
		ClientInterface: &adapters.ClientAdapter{
			Config: config.GetDataBaseConfig(),
		},
		ProductInterface: &adapters.ProductAdapter{
			Config: config.GetDataBaseConfig(),
		},
		RequestInterface: &adapters.RequestAdapter{
			Config: config.GetDataBaseConfig(),
		},
	}
	dbc.ClientInterface.Connect()
	dbc.ProductInterface.Connect()
	dbc.RequestInterface.Connect()
	return dbc
}
