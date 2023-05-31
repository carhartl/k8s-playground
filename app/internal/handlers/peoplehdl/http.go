package peoplehdl

import (
	"net/http"

	"github.com/carhartl/playground/internal/core/domain"
	"github.com/carhartl/playground/internal/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPHandler struct {
	peopleService ports.PeopleService
}

type personDTO struct {
	Email     string    `json:"email,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Uuid      uuid.UUID `json:"uuid,omitempty"`
}

func New(peopleService ports.PeopleService) HTTPHandler {
	return HTTPHandler{
		peopleService: peopleService,
	}
}

func (hdl HTTPHandler) Get(c *gin.Context) {
	person, err := hdl.peopleService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, personDTO(person))
}

func (hdl HTTPHandler) Create(c *gin.Context) {
	var pdto personDTO
	if err := c.ShouldBindJSON(&pdto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	person, err := hdl.peopleService.Create(domain.Person(pdto))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, personDTO(person))
}
