package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var loginFormTmpl = `
<html>
	<body>
	<form action="/get_cookie" method="post">
		Login: <input type="text" name="login">
		E-mail: <input type="email" name="Email">
		Password: <input type="password" name="Password">
		<input type="submit" value="Login">
	</form>
	</body>
</html>
`

var sessions = map[string]string{}

func Login(w http.ResponseWriter, r *http.Request) {

	sessionID, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		w.Write([]byte(loginFormTmpl))
		return
	} else if err != nil {
		PanicOnErr(err)
	}

	username, ok := sessions[sessionID.Value]

	if !ok {
		fmt.Fprint(w, "Session not found")
	} else {
		fmt.Fprint(w, "Welcome, "+username)
	}

}
func GetCookie(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	inputLogin := r.Form["login"][0]
	expiration := time.Now().Add(5 * time.Hour)

	sessionID := RandStringRunes(32)
	sessions[sessionID] = inputLogin

	cookie := http.Cookie{Name: "session_id", Value: sessionID, Expires: expiration}
	http.SetCookie(w, &cookie)

	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	lastInsertId := 0
	err = db.QueryRow(
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id_user",
		r.Form["login"][0],
		r.Form["Email"][0],
		r.Form["Password"][0],
	).Scan(&lastInsertId)

	fmt.Fprintf(w, "Insert - LastInsertId: %d \n", lastInsertId)

	//defer http.Redirect(w, r, "/login", http.StatusFound)

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
