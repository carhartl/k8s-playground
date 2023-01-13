package main

import (
	"fmt"
	"net/http"

	"github.com/carhartl/playground/internal/core/services/peoplesrv"
	"github.com/carhartl/playground/internal/handlers/peoplehdl"
	"github.com/carhartl/playground/internal/repositories/peoplerepo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "yb-tservers.yugabytedb-system"
	port     = 5433
	user     = "yugabyte"
	password = "yugabyte"
	dbname   = "yugabyte"
)

func main() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host= %s port = %d user = %s password = %s dbname = %s sslmode=disable",
		host, port, user, password, dbname))
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
