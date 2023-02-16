package peoplehdl

import (
	"net/http"
	"testing"

	"github.com/carhartl/playground/internal/core/domain"
	"github.com/carhartl/playground/internal/core/ports"
	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func init() {
	gin.SetMode(gin.TestMode)
}

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
	handler := New(srv)

	engine := gin.New()
	engine.GET("/people/:id", handler.Get)
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(engine),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	e.GET("/people/" + uuid).
		Expect().
		Status(http.StatusOK).
		JSON().Object().ContainsKey("uuid").ContainsValue(uuid)
}

func TestCreate(t *testing.T) {
	handler := New(new(mockService))

	engine := gin.New()
	engine.POST("/people", handler.Create)
	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(engine),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	e.POST("/people").
		Expect().
		Status(http.StatusBadRequest)

	e.POST("/people").WithBytes([]byte(`{"firstName":"foo"}`)).
		Expect().
		Status(http.StatusOK).
		JSON().Object().ContainsKey("firstName")

}
