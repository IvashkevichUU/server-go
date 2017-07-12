package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello World"
	})
	m.Get("/db", openDb)
	m.Get("/createdb", createDb)
	//m.Get("/createstudent", createStudent)
	//m.Get("/getstudents", getStudents)
	m.Get("/hello", HelloServer)

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

func createDb() {

	result, err := db.Exec("CREATE TABLE IF NOT EXISTS students (id SERIAL NOT NULL, fio CHARACTER VARYING(300) NOT NULL, info TEXT NOT NULL, score INTEGER NOT NULL )")
	if err != nil {
		log.Println(err)
	}
	affected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
	}
	fmt.Sprintf("Update - RowsAffected", affected)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет %s!\n", r.URL.Path[1:])
}
