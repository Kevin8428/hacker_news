package users

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/kevin8428/hackernews/domain"
)

type controller struct {
	Service
}

func makeController(us Service) controller {
	return controller{
		Service: us,
	}
}

func (c *controller) ShowUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		t, err := template.ParseFiles(wd + "/user_page.html")
		if err != nil {
			log.Fatal("error: ", err)
			return
		}
		r.ParseForm()
		userID := r.Form["id"][0]
		id, _ := strconv.Atoi(userID)
		user := c.Service.FindUser(id)
		u := domain.User{
			LastName: user.LastName,
		}

		err = t.Execute(w, u)
		if err != nil {
			log.Fatal("error: ", err)
			return
		}
	})
}

func (c *controller) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		saveUser := c.Service.SaveNewUser()
		fmt.Println("saved user: ", saveUser)
	})
}

func (c *controller) SaveUserArticle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.Form["name"][0]
		author := r.Form["author"][0]
		website := r.Form["website"][0]
		userID := r.Form["user_id"][0]
		id, _ := strconv.Atoi(userID)
		c.Service.SaveArticle(name, author, website, id)
	})
}