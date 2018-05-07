package authentication

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kevin8428/hackernews/repos"
	"golang.org/x/crypto/bcrypt"
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

func createUser(userRepo repos.UsersRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		fn := r.Form.Get("first-name")
		ln := r.Form.Get("last-name")
		token := GenerateToken(email)
		authToken, err := userRepo.CreateUser(token, email, password, fn, ln)
		if err != nil {
			fmt.Println("couldn't create user: ", err)
		} else if authToken != "" {
			json.NewEncoder(w).Encode(authToken)
		}
	})
}

func GenerateToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}
