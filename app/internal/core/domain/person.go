package domain

import "github.com/google/uuid"

//nolint:unused
type Person struct {
	Email     string
	FirstName string
	LastName  string
	Uuid      uuid.UUID
}
