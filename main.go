package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"text/template"
)

func renderPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "homepage.html")
}

type handler struct{}

type Person struct {
	Name   string
	Emails []string
}

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"Completed"`
	Due       time.Time `json:"Due"`
}
type Todos []Todo

type articles struct{}

func (h articles) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a := Todo{
		Name:      "kevin",
		Completed: false,
	}

	json.NewEncoder(w).Encode(a)
}

func getArticles() Todo {
	a := Todo{}
	res, err := http.Get("http://localhost:5050/articles")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(body, &a)
	if err != nil {
		fmt.Println("unmarshall error: ", err)
	}
	return a
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("homepage.html")
	person := Person{
		Name:   "kevin",
		Emails: []string{"kevin@mail.com", "deutscher@mail.com"},
	}
	articles := getArticles()
	fmt.Println("articles: ", articles)

	err = t.Execute(w, person)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}

func buildHandler(page *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page.Execute(w, r)
	})
}

func main() {
	////////////////implement without handler////////////////
	// http.HandleFunc("/", renderPage)
	// err := http.ListenAndServe(":5050", nil)
	// if err != nil {
	// 	panic(err)
	// }

	////////////////implement with handler////////////////
	server := http.NewServeMux()
	server.Handle("/articles", articles{})
	server.Handle("/", handler{})
	err := http.ListenAndServe(":5050", server)
	if err != nil {
		panic(err)
	}

	////////////////implement with handler function////////////////
	// server := http.NewServeMux()
	// homepage := template.Must(template.ParseFiles("homepage.html"))
	// server.Handle("/", buildHandler(homepage))
	// if err := http.ListenAndServe(":5050", server); err != nil {
	// 	panic(err)
	// }

}
