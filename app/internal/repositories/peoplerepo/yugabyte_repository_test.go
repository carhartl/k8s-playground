package peoplerepo

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/carhartl/playground/internal/core/domain"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "people" ("email","first_name","last_name","id") 
                                        VALUES ($1,$2,$3,$4) RETURNING "people"."id"`)).
		WithArgs("foo@example.com", "foo", "bar", "be2ae927-ad56-4e85-b75c-9101185c3b6b").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("be2ae927-ad56-4e85-b75c-9101185c3b6b"))
	mock.ExpectCommit()

	orm, err := gorm.Open("postgres", db)
	require.NoError(t, err)
	orm.LogMode(true)
	repo := New(orm)
	err = repo.Save(domain.Person{
		Email:     "foo@example.com",
		FirstName: "foo",
		LastName:  "bar",
		Uuid:      uuid.MustParse("be2ae927-ad56-4e85-b75c-9101185c3b6b"),
	})
	require.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestFindByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "people" WHERE (id::text = $1) ORDER BY "people"."id" ASC LIMIT 1`)).
		WithArgs("foo")

	orm, err := gorm.Open("postgres", db)
	require.NoError(t, err)
	orm.LogMode(true)
	repo := New(orm)
	_, err = repo.FindByID("foo")
	require.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
