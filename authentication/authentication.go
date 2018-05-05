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
		dbPassword, err := userRepo.GetPasswordUsingEmail(email)
		if err != nil {
			fmt.Println("couldn't get password")
		}
		fmt.Println("user input email: ", email)
		fmt.Println("user input password: ", password)
		fmt.Println("database password: ", dbPassword)
		type response struct {
			isValid bool
		}
		resp := response{false}
		if password == dbPassword {
			resp.isValid = true
		}
		json.NewEncoder(w).Encode(resp.isValid)
		// js, err := json.Marshal(resp)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// fmt.Println("resp: ", resp)
		// w.Header().Set("Content-Type", "application/json")
		// w.Write(js)
	})
}
