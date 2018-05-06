package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kevin8428/hackernews/repos"
)

func authenticateUser(userRepo repos.UsersRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		dbPassword, authToken, err := userRepo.GetPasswordUsingEmail(email)
		if err != nil {
			fmt.Println("couldn't get password")
		}
		if password == dbPassword {
			json.NewEncoder(w).Encode(authToken)
		}
	})
}
