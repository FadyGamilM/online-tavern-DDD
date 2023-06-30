package customer

import (
	"github.com/FadyGamilM/online-tavern-DDD/aggregates"
	"github.com/google/uuid"
)

// This repository is an interface for any data provider for the cutomer aggregate

// the errors of this aggregate are inside the aggregate package and the repo implementation wiil utilize them later
type ErrCustomerNotFound struct {
	ErrMsg string
}

func (e ErrCustomerNotFound) Error() string {
	return e.ErrMsg
}

type ErrFailedToAddCustomer struct {
	ErrMsg string
}

func (e ErrFailedToAddCustomer) Error() string {
	return e.ErrMsg
}

type ErrFailedToUpdateCustomer struct {
	ErrMsg string
}

func (e ErrFailedToUpdateCustomer) Error() string {
	return e.ErrMsg
}

type ErrFailedToDeleteCustomer struct {
	ErrMsg string
}

// -> interface abstraction
type CustomerRepo interface {
	Get(uuid.UUID) (aggregates.Customer, error)
	Add(aggregates.Customer) error
	Update(aggregates.Customer) error
}
