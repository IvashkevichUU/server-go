package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"github.com/lib/pq"

	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

var (
	db *sql.DB
)

type payment struct {
	Number  int                    `json:"number"`
	Success bool                   `json:"success"`
	Res     map[string]interface{} `json:"Res"`
}

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

	m.Get("/login", Login)
	m.Post("/get_cookie", GetCookie)

	m.Get("/createdbpayments", createDbPayment)
	m.Get("/createpayments", createPayment)

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

func createPayment(w http.ResponseWriter, r *http.Request) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	var id int
	row := db.QueryRow("SELECT id FROM public.payments ORDER BY id DESC LIMIT 1")
	err = row.Scan(&id)
	PanicOnErr(err)
	fmt.Fprintln(w, "PrintByID:", id)

	urlSend := "https://apibtc.com/api/create_wallet?token=4e71a0c5cbcf5004cc7977c32b6e917c79c5abda8f4aaceb456626d180f6771f&callback=https://woods.one/api/index.php?ID=" + strconv.Itoa(id)
	fmt.Fprintf(w, "Url Send: %s \n", urlSend)
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, urlSend, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	payment1 := payment{}
	jsonErr := json.Unmarshal(body, &payment1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Fprintln(w, "Status: ", payment1.Success)
	//fmt.Fprintln(w, payment1.Res["Sign"])
	//fmt.Fprintln(w, payment1.Res["Adress"])
	//fmt.Fprintln(w, payment1.Res["Address"])
	//fmt.Fprintln(w, payment1.Res["DataEnd"])

	lastInsertId := 0

	err = db.QueryRow(
		"INSERT INTO payments (address, amount) VALUES ($1, $2) RETURNING id",
		payment1.Res["Adress"],
		float64(id)*0.01,
	).Scan(&lastInsertId)

	fmt.Fprintf(w, "Insert - LastInsertId: %d , Payment address: %s , Amount: %v \n", lastInsertId, payment1.Res["Adress"], float64(id)*0.01)

}
