package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/go-faker/faker/v4"
)

type Person struct {
	Email     string `faker:"email"`
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
}

func main() {
	f, err := os.Create("people.csv")
	defer f.Close()
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()
	if err := w.Write([]string{"Email", "FirstName", "LastName"}); err != nil {
		log.Fatalln("Error writing record to file", err)
	}

	p := Person{}
	for i := 0; i < 500; i++ {
		err = faker.FakeData(&p)
		if err != nil {
			log.Fatalln("Error producing fake data", err)
		}
		row := []string{p.Email, p.FirstName, p.LastName}
		if err := w.Write(row); err != nil {
			log.Fatalln("Error writing record to file", err)
		}
	}
}
