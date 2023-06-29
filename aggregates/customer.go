package aggregates

import (
	"github.com/FadyGamilM/online-tavern-DDD/entities"
	"github.com/FadyGamilM/online-tavern-DDD/valueobjects"
)

// the aggregate contains the business logic, the entity is just a placeholder for info
type Customer struct {
	// Person is the root entity of the customer so the id of person is the pk of the customer aggregate
	person       *entities.Person
	drinks       []*entities.Drink_Item
	transactions []valueobjects.Transaction
}
