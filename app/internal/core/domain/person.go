package domain

import "github.com/google/uuid"

type Person struct {
	Email     string
	FirstName string
	LastName  string
	Uuid      uuid.UUID
}
