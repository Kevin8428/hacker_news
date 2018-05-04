package users

import (
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
		user := c.Service.FindUsersByUserID(id)
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
