package users

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"text/template"

	"github.com/kevin8428/hackernews/database"
)

// User struct
type User struct {
	id        int
	firstName string
	LastName  string
	Password  string
	Articles  string
}

func (u User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", database.Host, database.Port, database.User, database.Password, database.DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error1")
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	mydb := &database.DatabaseHandler{DB: db}
	r.ParseForm()
	userID := r.Form["id"][0]
	foo, _ := strconv.Atoi(userID)
	lastName := mydb.GetUserLastName(foo)
	article := mydb.GetUserFirstArticle(foo)
	user := User{
		LastName: lastName,
		Articles: article,
	}
	t, err := template.ParseFiles("user_page.html")
	err = t.Execute(w, user)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}

type AddArticle struct{}

func (a AddArticle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", database.Host, database.Port, database.User, database.Password, database.DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error1")
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	mydb := &database.DatabaseHandler{DB: db}
	r.ParseForm()
	name := r.Form["name"][0]
	author := r.Form["author"][0]
	website := r.Form["website"][0]
	userID := r.Form["user_id"][0]
	id, _ := strconv.Atoi(userID)
	mydb.AddArticle(name, author, website, id)
}
