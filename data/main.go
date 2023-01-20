package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/go-faker/faker/v4"
	"github.com/yugabyte/pgx/v4"
)

type Person struct {
	Email     string `faker:"email"`
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	var out bytes.Buffer
	csv := csv.NewWriter(&out)
	if err := csv.Write([]string{"email", "first_name", "last_name"}); err != nil {
		panic(err)
	}
	p := Person{}
	for i := 0; i < 500; i++ {
		err := faker.FakeData(&p)
		if err != nil {
			panic(err)
		}
		row := []string{p.Email, p.FirstName, p.LastName}
		if err := csv.Write(row); err != nil {
			panic(err)
		}
	}
	csv.Flush()

	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		getEnv("PGUSER", "postgres"),
		getEnv("PGPASSWORD", "postgres"),
		getEnv("PGHOST", "localhost"),
		getEnv("PGPORT", "5432"),
		getEnv("PGDATABASE", "postgres")))
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close(context.Background()) }()

	_, err = conn.Exec(context.Background(), `
		CREATE EXTENSION IF NOT EXISTS pgcrypto;
		CREATE TABLE IF NOT EXISTS people (
			id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			email text,
			first_name text,
			last_name text
		);
	`)
	if err != nil {
		panic(err)
	}

	res, err := conn.PgConn().CopyFrom(context.Background(), bytes.NewReader(out.Bytes()), "COPY people(email,first_name,last_name) FROM STDIN (FORMAT csv, HEADER true)")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Database populated with %d records", res.RowsAffected())
}
