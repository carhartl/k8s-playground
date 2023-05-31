package peoplerepo

import (
	"errors"

	"github.com/carhartl/playground/internal/core/domain"
	"github.com/carhartl/playground/internal/core/ports"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type person struct {
	Email     string    `gorm:"column:email"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Uuid      uuid.UUID `gorm:"column:id;type:uuid;primary_key"`
}

var (
	// ErrNotFound is a convenience reference for the actual GORM error
	ErrNotFound = gorm.ErrRecordNotFound
)

type yugabyteRepository struct {
	ports.PeopleRepository
	db *gorm.DB
}

func New(db *gorm.DB) ports.PeopleRepository {
	return yugabyteRepository{
		db: db,
	}
}

func (repo yugabyteRepository) Save(newPerson domain.Person) error {
	new := person(newPerson)
	res := repo.db.Create(&new)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo yugabyteRepository) FindByID(id string) (domain.Person, error) {
	var person person
	res := repo.db.Where("id::text = ?", id).First(&person)
	if errors.Is(res.Error, ErrNotFound) {
		return domain.Person{}, ErrNotFound
	}
	return domain.Person{
		Email:     person.Email,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Uuid:      person.Uuid,
	}, nil
}
