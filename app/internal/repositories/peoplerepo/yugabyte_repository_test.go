package peoplerepo

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/carhartl/playground/internal/core/domain"
	"github.com/carhartl/playground/internal/core/ports"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo ports.PeopleRepository
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	orm, err := gorm.Open("postgres", db)
	require.NoError(s.T(), err)
	orm.LogMode(true)

	s.repo = New(orm)
}

func (s *Suite) TestSave() {
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "people" ("email","first_name","last_name","id") 
                                        VALUES ($1,$2,$3,$4) RETURNING "people"."id"`)).
		WithArgs("foo@example.com", "foo", "bar", "be2ae927-ad56-4e85-b75c-9101185c3b6b").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("be2ae927-ad56-4e85-b75c-9101185c3b6b"))
	s.mock.ExpectCommit()

	err := s.repo.Save(domain.Person{
		Email:     "foo@example.com",
		FirstName: "foo",
		LastName:  "bar",
		Uuid:      uuid.MustParse("be2ae927-ad56-4e85-b75c-9101185c3b6b"),
	})
	require.NoError(s.T(), err)

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.T().Errorf("There were unfulfilled expectations: %s", err)
	}
}

func (s *Suite) TestFindByID() {
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "people" WHERE (id::text = $1) ORDER BY "people"."id" ASC LIMIT 1`)).
		WithArgs("foo")

	_, err := s.repo.FindByID("foo")
	require.NoError(s.T(), err)

	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.T().Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
