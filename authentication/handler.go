package authentication

import (
	"net/http"

	"github.com/kevin8428/hackernews/repos"
)

func InitializeHandler(server *http.ServeMux, userRepo repos.UsersRepository) {
	server.Handle("/sign-in", authenticateUser(userRepo))
	server.Handle("/sign-up", createUser(userRepo))
}
