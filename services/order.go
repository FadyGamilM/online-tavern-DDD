package services

import (
	"github.com/FadyGamilM/online-tavern-DDD/domain/customer"
	"github.com/FadyGamilM/online-tavern-DDD/domain/customer/memory"
)

// i will utilize the service generator configuration pattern for more abstraction and more modularity

type OrderService struct {
	// the OrderService will need to deal with the customer aggregate, thats why we have the customer repository here because its the manager of the aggregate
	customers_repo customer.CustomerRepo
}

// the OrderConfiguration is an alias of a function signature type
// => takes a pointer to OrderService and returns an error or nil if there is no error
// HINT : i recieve a pointer to the service because i will modify the service based on the configurations
type OrderConfiguration func(orderService *OrderService) error

// Factory method to create an order service and apply configurations to it before creation
func NewOrderService(configs ...OrderConfiguration) (*OrderService, error) {
	// loop through all the configurations and apply them to the OrderService
	// 1. create a ref for the order Service
	orderService := &OrderService{}

	for _, cfg := range configs {
		// 2. apply the cfg to the order service
		err := cfg(orderService)
		if err != nil {
			return nil, err
		}
	}
	return orderService, nil
}

// method which returns a function that matches the order configurations type
// HINT => this method accepts the repo interface so it accepts any type of data provider
func WithCustomerRepo(customerRepo customer.CustomerRepo) OrderConfiguration {
	// when we call the cfg(orderService) above, so we call the following methods which apply the repo to the service
	return func(orderService *OrderService) error {
		orderService.customers_repo = customerRepo
		return nil
	}
}

func WithMemoryCustomerRepo() OrderConfiguration {
	repo := memory.New()
	return WithCustomerRepo(repo)
}

// now we can have a method for WithMongoCustomerRepo and WithMySqlCustomerRepo
// then create the orderService as following : services.NewOrderService(WithMongoCustomerRepo())
// if we need to chagne we can replace the WithMongo to WithMySql .. simple ha
