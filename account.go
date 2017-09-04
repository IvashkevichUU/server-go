package main

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//var loginFormTmpl = `
//<html>
//	<body>
//	<form action="/auth" method="post">
//		Login: <input type="text" name="login">
//		Password: <input type="password" name="Password">
//		<input type="submit" value="Login">
//	</form>
//	</body>
//</html>
//`

//var RegFormTmpl = `
//<html>
//	<body>
//	<form action="/get_cookie" method="post">
//		Login: <input type="text" name="login">
//		E-mail: <input type="email" name="Email">
//		Password: <input type="password" name="Password">
//		<input type="submit" value="Login">
//	</form>
//	</body>
//</html>
//`

var sessions = map[string]string{}

type Person struct {
	Name   string
	Return []byte
}

func Register(w http.ResponseWriter, r *http.Request) {

	sessionID, err := r.Cookie("session_id")

	if err == http.ErrNoCookie || sessions[sessionID.Value] == "" {
		t, _ := template.ParseFiles("templates/registration.html")
		t.Execute(w, "active")
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

func Login(w http.ResponseWriter, r *http.Request) {

	sessionID, err := r.Cookie("session_id")

	if err == http.ErrNoCookie || sessions[sessionID.Value] == "" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, "active")
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

func Auth(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	inputLogin := r.Form["login"][0]
	inputPass := r.Form["Password"][0]

	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println(err)
	}

	var name string
	var password string
	row := db.QueryRow("SELECT name, password FROM users WHERE name = $1", inputLogin)

	err = row.Scan(&name, &password)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	if inputPass == password {
		expiration := time.Now().Add(5 * time.Hour)

		sessionID := RandStringRunes(32)
		sessions[sessionID] = inputLogin

		cookie := http.Cookie{Name: "session_id", Value: sessionID, Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/account", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
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

	http.Redirect(w, r, "/account", http.StatusFound)

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Accounts(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	} else if err != nil {
		PanicOnErr(err)
	}

	username, ok := sessions[sessionID.Value]

	if !ok {
		http.Redirect(w, r, "/logout", http.StatusFound)
		return
	} else {

		p := Person{}
		p.Name = username
		//p.Return = Websocket(username)
		t, _ := template.ParseFiles("templates/account.html")
		t.Execute(w, p)
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	sessions[sessionID.Value] = ""

	http.Redirect(w, r, "/", 302)
}

//
//func Websocket(account string) []byte {
//	origin := "http://localhost/"
//	url := "wss://bitshares.openledger.info/ws"
//	ws, err := websocket.Dial(url, "", origin)
//	if err != nil {
//		log.Fatal(err)
//	}
//	per := "{\"id\": 1, \"method\": \"call\", \"params\": [1, \"login\", [\"\", \"\"]]}"
//	per2 := "{\"id\": 2, \"method\": \"call\", \"params\": [1,\"database\",[]]}"
//	per3 := "{\"id\": 4, \"method\": \"call\", \"params\": [2,\"get_full_accounts\",[[" + account + "], false]]}"
//	if _, err := ws.Write([]byte(per)); err != nil {
//		log.Fatal(err)
//	}
//	var msg = make([]byte, 2048)
//	var n int
//	if n, err = ws.Read(msg); err != nil {
//		log.Fatal(err)
//	}
//	mes1 := msg[:n]
//
//	if _, err := ws.Write([]byte(per2)); err != nil {
//		log.Fatal(err)
//	}
//
//	if n, err = ws.Read(msg); err != nil {
//		log.Fatal(err)
//	}
//	mes2 := msg[:n]
//
//	if _, err := ws.Write([]byte(per3)); err != nil {
//		log.Fatal(err)
//	}
//
//	if n, err = ws.Read(msg); err != nil {
//		log.Fatal(err)
//	}
//	mes3 := msg[:n]
//	fmt.Println(mes1, mes2)
//	return mes3
//
//}
