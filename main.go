package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alfg/blockchain"
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
	m.Get("/createstudent", createStudent)
	m.Get("/getstudents", getStudents)
	m.Get("/getstudents/{id}", PrintByID)
	m.Get("/hello", HelloServer)
	m.Get("/blockchain", Blockchains)
	m.Get("/createdbpayments", createDbPayment)

	m.Run()
}

func Blockchains(w http.ResponseWriter, r *http.Request) {

	c, err := blockchain.New()
	resp, err := c.GetAddress("1AJbsFZ64EpEfS5UAjAfcUG8pH8Jn3rn1F")

	if err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "Hash160: %s, Address: %s, NTx: %d, TotalReceived: %d, TotalSent: %d, FinalBalance: %d\n", resp.Hash160, resp.Address, resp.NTx, resp.TotalReceived, resp.TotalSent, resp.FinalBalance)
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

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintByID(id int64) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	var fio string
	var info sql.NullString
	// var info string
	var score int
	row := db.QueryRow("SELECT fio, info, score FROM students WHERE id = $1", id)
	// fmt.Println(row)
	err = row.Scan(&fio, &info, &score)
	PanicOnErr(err)
	fmt.Println("PrintByID:", id, "fio:", fio, "info:", info, "score:", score)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	lastInsertId := 0
	err = db.QueryRow(
		"INSERT INTO students (fio, info, score) VALUES ($1, $2, $3) RETURNING id",
		"Oleg Petrov",
		"test student",
		"85",
	).Scan(&lastInsertId)

	fmt.Fprintf(w, "Insert - LastInsertId: %d \n", lastInsertId)

	PrintByID(int64(lastInsertId))

}

func getStudents(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	err = db.Ping()
	PanicOnErr(err)

	rows, err := db.Query("SELECT * FROM students")
	PanicOnErr(err)
	for rows.Next() {
		var id uint
		var fio string
		var info string
		var score uint
		err = rows.Scan(&id, &fio, &info, &score)
		PanicOnErr(err)

		fmt.Fprintf(w, "id: %d, Fio: %s, Info: %s, Score: %d\n", id, fio, info, score)
	}
	fmt.Fprintln(w, "Open connections: ", db.Stats().OpenConnections)
	rows.Close()
}

func createDbPayment() {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	result, err := db.Exec("CREATE TABLE IF NOT EXISTS payments (id SERIAL NOT NULL, address CHARACTER VARYING(300) NOT NULL, amount FLOAT NOT NULL )")
	if err != nil {
		log.Println(err)
	}
	affected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
	}
	fmt.Sprintf("Update - RowsAffected", affected)
}
