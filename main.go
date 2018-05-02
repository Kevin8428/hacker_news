package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/kevin8428/hackernews/api"
	"github.com/kevin8428/hackernews/users"
	_ "github.com/lib/pq"
)

func renderPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "homepage.html")
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("homepage.html")
	articles := api.GetArticles()

	err = t.Execute(w, articles)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}

func main() {

	////////////////implement with handler////////////////
	server := http.NewServeMux()
	server.Handle("/articles", api.Articles{})
	server.Handle("/user", users.User{})
	server.Handle("/add-article", users.AddArticle{})
	server.Handle("/", handler{})
	err := http.ListenAndServe(":5050", server)
	if err != nil {
		panic(err)
	}

	////////////////implement without handler////////////////
	// http.HandleFunc("/", renderPage)
	// err := http.ListenAndServe(":5050", nil)
	// if err != nil {
	// 	panic(err)
	// }

	////////////////implement with handler function////////////////
	// server := http.NewServeMux()
	// homepage := template.Must(template.ParseFiles("homepage.html"))
	// server.Handle("/", buildHandler(homepage))
	// if err := http.ListenAndServe(":5050", server); err != nil {
	// 	panic(err)
	// }

}
