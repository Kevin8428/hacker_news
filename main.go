package main

import (
	"net/http"

	"github.com/kevin8428/hackernews/api"
	"github.com/kevin8428/hackernews/articles"
	"github.com/kevin8428/hackernews/authentication"
	"github.com/kevin8428/hackernews/repos"
	"github.com/kevin8428/hackernews/users"
	"github.com/kevin8428/hackernews/websockets"
	_ "github.com/lib/pq"
)

func main() {
	database := repos.Initialize()
	defer database.Articles.DB.Close()
	defer database.Users.DB.Close()
	as := articles.NewService(database.Articles)
	us := users.NewService(database.Users)
	h := websockets.NewHub()
	server := http.NewServeMux()
	server.Handle("/static/app.js", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/app.js")
	}))
	server.Handle("/static/style.css", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/style.css")
	}))
	articles.InitializeHandler(server, as)
	users.InitializeHandler(server, us)
	authentication.InitializeHandler(server, *database.Users)
	api.InitializeHandler(server)
	websockets.InitializeHandler(server, h)
	err := http.ListenAndServe(":5050", server)
	if err != nil {
		panic(err)
	}
}
