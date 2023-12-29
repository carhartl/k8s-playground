package peoplehdl

import (
	"net/http"

	"github.com/carhartl/playground/internal/core/domain"
	"github.com/carhartl/playground/internal/core/ports"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type personDTO struct {
	Email     string    `json:"email,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Uuid      uuid.UUID `json:"uuid,omitempty"`
}

type HTTPHandler struct {
	peopleService ports.PeopleService
}

func New(srv ports.PeopleService) HTTPHandler {
	return HTTPHandler{
		peopleService: srv,
	}
}

func (hdl HTTPHandler) Get(c echo.Context) error {
	person, err := hdl.peopleService.Get(c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, personDTO(person))
}

func (hdl HTTPHandler) Create(c echo.Context) error {
	var pdto personDTO
	if err := c.Bind(&pdto); err != nil {
		return err
	}
	person, err := hdl.peopleService.Create(domain.Person(pdto))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, personDTO(person))
}
