package ports

import "github.com/carhartl/playground/internal/core/domain"

// Driven port
type PeopleRepository interface {
	Save(person domain.Person) error
	FindByID(personID string) (domain.Person, error)
}
