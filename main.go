package main

import (
	"net/http"

	"github.com/kevin8428/hackernews/api"
	"github.com/kevin8428/hackernews/articles"
	"github.com/kevin8428/hackernews/repos"
	"github.com/kevin8428/hackernews/users"
	_ "github.com/lib/pq"
)

func main() {
	database := repos.Initialize()
	defer database.Articles.DB.Close()
	defer database.Users.DB.Close()
	as := articles.NewService(database.Articles)
	us := users.NewService(database.Users)
	server := http.NewServeMux()
	articles.InitializeHandler(server, as)
	users.InitializeHandler(server, us)
	server.Handle("/articles", api.Articles{})
	err := http.ListenAndServe(":5050", server)
	if err != nil {
		panic(err)
	}
}
