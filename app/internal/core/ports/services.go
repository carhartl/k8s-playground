package ports

import "github.com/carhartl/playground/internal/core/domain"

// Driving port (use cases)
type PeopleService interface {
	Get(id string) (domain.Person, error)
	Create(person domain.Person) (domain.Person, error)
}
