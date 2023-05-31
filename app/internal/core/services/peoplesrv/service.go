package peoplesrv

import (
	"errors"

	"github.com/carhartl/playground/internal/core/domain"
	"github.com/carhartl/playground/internal/core/ports"
	"github.com/google/uuid"
)

type service struct {
	peopleRepository ports.PeopleRepository
}

func New(repo ports.PeopleRepository) service {
	return service{
		peopleRepository: repo,
	}
}

func (srv service) Get(id string) (domain.Person, error) {
	person, err := srv.peopleRepository.FindByID(id)
	if err != nil {
		return domain.Person{}, errors.New("Finding person in repository has failed")
	}
	return person, nil
}

func (srv service) Create(person domain.Person) (domain.Person, error) {
	person.Uuid = uuid.New()
	if err := srv.peopleRepository.Save(person); err != nil {
		return domain.Person{}, errors.New("Creating person in repository has failed")
	}
	return person, nil
}
