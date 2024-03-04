package containers

type AppContainer struct {
	DataBaseContainer DataBaseContainer
	ClientContainer   ClientContainer
	ProductContainer  ProductContainer
	RequestContainer  RequestContainer
	ServerContainer   ServerContainer
}

func NewAppContainer() AppContainer {
	var dataBaseContainer *DataBaseContainer = NewDataBaseContainer()
	var clientContainer ClientContainer = NewClientContainer(dataBaseContainer.ClientInterface)
	var productContainer ProductContainer = NewProductContainer(dataBaseContainer.ProductInterface)
	var requestContainer RequestContainer = NewRequestContainer(dataBaseContainer.RequestInterface)
	var serverContainer ServerContainer = NewServerContainer(clientContainer.Router, productContainer.Router, requestContainer.Router)
	return AppContainer{
		DataBaseContainer: *dataBaseContainer,
		ClientContainer:   clientContainer,
		ProductContainer:  productContainer,
		RequestContainer:  requestContainer,
		ServerContainer:   serverContainer,
	}
}
