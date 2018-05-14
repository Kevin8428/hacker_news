package authentication

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/kevin8428/hackernews/repos"
	"golang.org/x/crypto/bcrypt"
)

type UserResponse struct {
	LastName  string `json:"family_name"`
	Email     string `json:"email"`
	FirstName string `json:"given_name"`
	Verified  string `json:"email_verified"`
}

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
			fmt.Println("couldn't get password: ", err)
		}
		err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
		// if password == dbPassword {
		// 	json.NewEncoder(w).Encode(authToken)
		// }
		if err == nil {
			json.NewEncoder(w).Encode(authToken)
		}
	})
}

func authenticateOAuth(userRepo repos.UsersRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UR := UserResponse{}
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		idToken := r.Form.Get("idtoken")
		resp, err := http.Get("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=" + idToken)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &UR)
		if err != nil {
			fmt.Println("error parsing json: ", err)
			panic(err)
		}
		_, authToken, err := userRepo.GetPasswordUsingEmail(UR.Email)
		if err == sql.ErrNoRows {
			w.Write(json.RawMessage("EOF"))
		} else if authToken != "" && UR.Verified == "true" {
			// createUserFromOAuth(UR)
			json.NewEncoder(w).Encode(authToken)
		} else {
			panic(err)
		}
	})
}

func createUserFromOAuth(ur UserResponse) error {
	_, err := http.PostForm(
		"/sign-up",
		url.Values{
			"email":      {ur.Email},
			"password":   {""},
			"first-name": {ur.FirstName},
			"last-name":  {ur.LastName},
			"is-oauth":   {"true"},
		})
	return err
	// Code to process response (written in Get request snippet) goes here
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
		// ln := r.Form.Get("is-oauth")
		token := GenerateToken(email)
		encryptedPassword, err := hashPassword(password)
		if err != nil {
			panic("problem encrypting password")
		}
		authToken, err := userRepo.CreateUser(token, email, encryptedPassword, fn, ln)
		if err != nil {
			fmt.Println("couldn't create user: ", err)
		} else if authToken != "" {
			json.NewEncoder(w).Encode(authToken)
		}
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
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
