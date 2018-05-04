package users

import (
	"net/http"
)

func InitializeHandler(server *http.ServeMux, us Service) {
	controller := makeController(us)
	server.Handle("/user", controller.ShowUser())
	server.Handle("/create-user", controller.CreateUser())
	server.Handle("/save-article", controller.SaveUserArticle())
}
