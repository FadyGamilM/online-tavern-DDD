package valueobjects

import "github.com/google/uuid"

// the feilds of the Transaction struct are in lowercase because i won't allow any other domain to modify the valueobject once its created
type Transaction struct {
	amount     int       `json:"amount"`
	from       uuid.UUID `json:"from"`
	to         uuid.UUID `json:"to"`
	created_at int       `json:"created_at"`
}
