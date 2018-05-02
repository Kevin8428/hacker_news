package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kevin8428/hackernews/api"

	"text/template"
)

func renderPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "homepage.html")
}

type handler struct{}
type articles struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("homepage.html")
	articles := api.GetArticles()
	fmt.Println("articles: ", articles)

	err = t.Execute(w, articles)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}

func (h articles) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a := []api.Article{
		{
			Name:   "article 1",
			Author: "kevin",
		},
		{
			Name:   "article 2",
			Author: "matt",
		},
		{
			Name:   "article 3",
			Author: "dave",
		},
		{
			Name:   "article 4",
			Author: "ben",
		},
	}

	json.NewEncoder(w).Encode(a)
}

func main() {
	////////////////implement with handler////////////////
	server := http.NewServeMux()
	server.Handle("/articles", articles{})
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

// func buildHandler(page *template.Template) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		page.Execute(w, r)
// 	})
// }
