package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-martini/martini"
	"github.com/lib/pq"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World"
	})
	m.Get("/db", openDb)
	m.Run()
}

func openDb() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	return db
}
