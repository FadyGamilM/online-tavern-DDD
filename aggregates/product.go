package aggregates

import (
	"github.com/FadyGamilM/online-tavern-DDD/entities"
	"github.com/FadyGamilM/online-tavern-DDD/helpers"
	"github.com/google/uuid"
)

type Product struct {
	item  *entities.Drink_Item
	price float64
	qty   int
}

func NewProduct(itemName, itemDescp string, price float64) (Product, error) {
	if itemDescp == "" || itemName == "" || price <= 0.0 {
		return Product{}, helpers.ErrMissingEssentialData{ErrMsg: "missing essential data fields to create a product"}
	}

	return Product{
		item:  &entities.Drink_Item{ID: uuid.New(), DrinkName: itemName, Description: itemDescp},
		price: price,
		qty:   1,
	}, nil
}

func (p *Product) UpdateQty(qty int) {
	p.qty = qty
}

func (p *Product) UpdatePrice(price float64) {
	p.price = price
}

// we can't get the ID of the product aggregate from the item entity because we only deal with aggregate so we have to implement this method
func (p *Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p *Product) GetItem() *entities.Drink_Item {
	return p.item
}
