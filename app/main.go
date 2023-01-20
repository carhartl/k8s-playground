package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/carhartl/playground/internal/core/services/peoplesrv"
	"github.com/carhartl/playground/internal/handlers/peoplehdl"
	"github.com/carhartl/playground/internal/repositories/peoplerepo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	router := gin.Default()

	router.GET("/people/:id", hdl.Get)
	router.POST("/people", hdl.Create)
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Status": "UP"})
	})

	fmt.Println("Running people service. Press Ctrl+C to exit...")
	if err := router.Run(":8888"); err != nil {
		fmt.Println("Error starting people server...")
	}
}
