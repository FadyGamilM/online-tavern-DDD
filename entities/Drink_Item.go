package entities

import "github.com/google/uuid"

type Drink_Item struct {
	ID          uuid.UUID `json:"id"`
	DrinkName   string    `json:"drink_names"`
	Description string    `json:"description"`
}
