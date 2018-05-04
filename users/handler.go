package users

import (
	"net/http"
)

func InitializeHandler(server *http.ServeMux, us Service) {
	controller := makeController(us)
	server.Handle("/user", controller.ShowUser())
}
