package entities

import "github.com/google/uuid"

// Person is a type which represents the Person entity inside all sub-domains
type Person struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}
