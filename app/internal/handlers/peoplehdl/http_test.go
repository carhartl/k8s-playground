package peoplehdl

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/carhartl/playground/internal/core/domain"
	"github.com/carhartl/playground/internal/core/ports"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	ports.PeopleService
	mock.Mock
}

func (m *mockService) Get(id string) (domain.Person, error) {
	args := m.Called(id)
	return domain.Person{Uuid: uuid.MustParse(args.String(0))}, args.Error(1)
}

func (m *mockService) Create(person domain.Person) (domain.Person, error) {
	return person, nil
}

func TestGet(t *testing.T) {
	const uuid = "3d298c59-ff91-4b43-b00c-378511474275"

	srv := new(mockService)
	srv.On("Get", uuid).Return(uuid, nil)
	hdl := New(srv)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/people/:id")
	c.SetParamNames("id")
	c.SetParamValues(uuid)

	if assert.NoError(t, hdl.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"uuid":"3d298c59-ff91-4b43-b00c-378511474275"}`, rec.Body.String())
	}
}

func TestCreate(t *testing.T) {
	hdl := New(new(mockService))
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"firstName":"foo"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, hdl.Create(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// The uuid below is the zero value due to mocking the service...
		assert.JSONEq(t, `{"uuid":"00000000-0000-0000-0000-000000000000", "firstName":"foo"}`, rec.Body.String())
	}
}
