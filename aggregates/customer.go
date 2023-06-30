package aggregates

import (
	"github.com/FadyGamilM/online-tavern-DDD/entities"
	"github.com/FadyGamilM/online-tavern-DDD/helpers"
	"github.com/FadyGamilM/online-tavern-DDD/valueobjects"
	"github.com/google/uuid"
)

// the aggregate contains the business logic, the entity is just a placeholder for info
type Customer struct {
	// Person is the root entity of the customer so the id of person is the pk of the customer aggregate
	person       *entities.Person
	drinks       []*entities.Drink_Item
	transactions []valueobjects.Transaction
}

// Factory Method Pattern to create a new aggregate
func NewCustomer(customer_name string, customer_age int) (Customer, error) {
	if customer_name == "" {
		return Customer{}, helpers.ErrInvalidPersonName{ErrMsg: "customer must have a valid name"}
	}

	if customer_age < 0 {
		return Customer{}, helpers.ErrInvalidPersonAge{ErrMsg: "customer must have a valid age"}
	}

	// the root entity of the customer aggregate
	person := &entities.Person{ID: uuid.New(), Name: customer_name, Age: customer_age}

	// the other entities and value objects
	drinks := make([]*entities.Drink_Item, 0)
	transactions := make([]valueobjects.Transaction, 0)

	return Customer{person: person, drinks: drinks, transactions: transactions}, nil

}

// exposed methods to deal with the customer aggregate without exposing its internal impl
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) GetAge() int {
	return c.person.Age
}
