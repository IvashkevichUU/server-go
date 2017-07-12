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
	m.Get("/createstudent", createStudent)
	m.Get("/getstudents", getStudents)
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

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

// PrintByID print student by id
func PrintByID(id int64) {
	var fio string
	var info sql.NullString
	// var info string
	var score int
	row := db.QueryRow("SELECT fio, info, score FROM students WHERE id = $1", id)
	// fmt.Println(row)
	err := row.Scan(&fio, &info, &score)
	PanicOnErr(err)
	fmt.Println("PrintByID:", id, "fio:", fio, "info:", info, "score:", score)
}

func createStudent() {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	// Exec исполняет запрос и возвращает сколько строк было затронуто и последнйи ИД вставленной записи
	// символ ? является placeholder-ом. все последующие значения авто-экранируются и подставляются с правильным кавычками
	var lastID int64
	err = db.QueryRow(
		"INSERT INTO students (fio, info, score) VALUES ($1, $2, $3) RETURNING id",
		"Ivan Ivanov",
		"info studenta",
		"23",
	).Scan(&lastID)
	PanicOnErr(err)

	fmt.Println("Insert - LastInsertId: ", lastID)

	PrintByID(lastID)
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	// проверяем что подключение реально произошло ( делаем запрос )
	err = db.Ping()
	PanicOnErr(err)

	// итерируемся по многим записям
	// Exec исполняет запрос и возвращает записи
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
	rows.Close()
}
