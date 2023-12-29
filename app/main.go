package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/carhartl/playground/internal/core/services/peoplesrv"
	"github.com/carhartl/playground/internal/handlers/peoplehdl"
	"github.com/carhartl/playground/internal/repositories/peoplerepo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host= %s port = %s user = %s password = %s dbname = %s sslmode=disable",
		getEnv("PGHOST", "localhost"),
		getEnv("PGPORT", "5432"),
		getEnv("PGUSER", "postgres"),
		getEnv("PGPASSWORD", "postgres"),
		getEnv("PGDATABASE", "postgres")))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	defer db.Close()

	repo := peoplerepo.New(db)
	srv := peoplesrv.New(repo)
	hdl := peoplehdl.New(srv)

	e := echo.New()
	e.GET("/people/:id", hdl.Get)
	e.POST("/people", hdl.Create)
	e.GET("/healthz", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":8888"))
}
