package product

import (
	"github.com/FadyGamilM/online-tavern-DDD/aggregates"
	"github.com/google/uuid"
)

// manage the product aggregate
type ProductRepository interface {
	GetAll() ([]aggregates.Product, error)
	Get(uuid.UUID) (aggregates.Product, error)
	Create(aggregates.Product) error
	Update(aggregates.Product) error
}
