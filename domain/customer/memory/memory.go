package memory

import (
	"sync"

	"github.com/FadyGamilM/online-tavern-DDD/aggregates"
	"github.com/FadyGamilM/online-tavern-DDD/domain/customer"
	"github.com/google/uuid"
)

// define the In-Memory data structre to store data
type MemoryRepo struct {
	// keep the internals of the data storage as package-access
	customers map[uuid.UUID]aggregates.Customer
	// sync with mutex
	sync.Mutex
}

// expose factory method on this data structre
func New() *MemoryRepo {
	return &MemoryRepo{
		customers: make(map[uuid.UUID]aggregates.Customer),
	}
}

// implemen the interface of repo via this data structure
func (repo *MemoryRepo) Get(id uuid.UUID) (aggregates.Customer, error) {
	if customer, ok := repo.customers[id]; ok {
		return customer, nil
	}
	return aggregates.Customer{}, customer.ErrCustomerNotFound{ErrMsg: "cannot find a customer with this id"}
}

func (repo *MemoryRepo) Add(c aggregates.Customer) error {
	// 1. check if the structure already exists [if there is a table in the database with tihs name :D]
	if repo.customers == nil {
		// lock for safety
		repo.Lock()
		repo.customers = make(map[uuid.UUID]aggregates.Customer)
		repo.Unlock()
	}

	// 2. check if there is a customer with this id exists in database before
	if _, ok := repo.customers[c.GetID()]; ok {
		return customer.ErrFailedToAddCustomer{ErrMsg: "customer already exists"}
	}

	// 3. create new customer and add it
	repo.Lock()
	repo.customers[c.GetID()] = c
	repo.Unlock()
	return nil
}

func (repo *MemoryRepo) Update(c aggregates.Customer) error {
	// 1. check if the structure already exists [if there is a table in the database with tihs name :D]
	if repo.customers == nil {
		// lock for safety
		repo.Lock()
		repo.customers = make(map[uuid.UUID]aggregates.Customer)
		repo.Unlock()
	}

	// 2. check if there is a customer with this id exists in database before
	if _, ok := repo.customers[c.GetID()]; !ok {
		return customer.ErrFailedToUpdateCustomer{ErrMsg: "customer doesn't exist"}
	}

	// 3. update the customer
	repo.Lock()
	repo.customers[c.GetID()] = c
	repo.Unlock()
	return nil
}
