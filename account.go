package main

import (
	"database/sql"
	"encoding/json"
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
	Name                   string
	Return                 interface{}
	Id                     string
	Memo_key               string
	Lifetime_referrer_name string
	Referrer_name          string
	Registrar_name         string
	Asset_type             string
	Balance                int
	CoinId                 string
	Owner                  string
}

type Response struct {
	Id      int               `json:"id"`
	Jsonrpc string            `json:"jsonrpc"`
	Result  [][]ResultTwoType `json:"result"`
}
type ResultOneType struct {
	_             string `json:".result[0][0]"`
	ResultTwoType `json:"result[0][1]"`
}
type ResultTwoType struct {
	Account AccountType `json:"account"`
	//Assets                 []string    `json:"assets"`
	Lifetime_referrer_name string `json:"lifetime_referrer_name"`
	Referrer_name          string `json:"referrer_name"`
	Registrar_name         string `json:"registrar_name"`
	//Statistics             interface{} `json:"statistics"`
	Balances []BalancesType `json:"balances"`
}
type BalancesType struct {
	Asset_type string `json:"asset_type"`
	Balance    int    `json:"balance"`
	Id         string `json:"id"`
	Owner      string `json:"owner"`
}
type AccountType struct {
	Id      string      `json:"id"`
	Options OptionsType `json:"options"`
}
type OptionsType struct {
	Memo_key string `json:"memo_key"`
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

		//type test interface{}
		//var itest test
		var itest Response
		var jtest = Websocket(username)
		err := json.Unmarshal(jtest, &itest)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(itest)

		p := Person{}
		p.Name = username
		p.Return = itest
		p.Id = itest.Result[0][1].Account.Id
		p.Registrar_name = itest.Result[0][1].Registrar_name
		p.Referrer_name = itest.Result[0][1].Referrer_name
		p.Lifetime_referrer_name = itest.Result[0][1].Lifetime_referrer_name
		p.Memo_key = itest.Result[0][1].Account.Options.Memo_key
		if itest.Result[0][1].Balances != nil {
			p.Asset_type = itest.Result[0][1].Balances[0].Asset_type
			p.Balance = itest.Result[0][1].Balances[0].Balance
			p.CoinId = itest.Result[0][1].Balances[0].Id
			p.Owner = itest.Result[0][1].Balances[0].Owner
		}
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

func Websocket(account string) []byte {
	origin := "http://localhost/"
	url := "wss://bitshares.openledger.info/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	per := "{\"id\": 1, \"method\": \"call\", \"params\": [1, \"login\", [\"\", \"\"]]}"
	per2 := "{\"id\": 2, \"method\": \"call\", \"params\": [1,\"database\",[]]}"
	per3 := "{\"id\": 4, \"method\": \"call\", \"params\": [2,\"get_full_accounts\",[[\"" + account + "\"], false]]}"
	if _, err := ws.Write([]byte(per)); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 2048)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	mes1 := msg[:n]

	if _, err := ws.Write([]byte(per2)); err != nil {
		log.Fatal(err)
	}

	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	mes2 := msg[:n]

	if _, err := ws.Write([]byte(per3)); err != nil {
		log.Fatal(err)
	}

	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	mes3 := msg[:n]
	fmt.Println(mes1, mes2)

	return mes3

}
