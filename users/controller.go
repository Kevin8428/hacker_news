package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		articles := c.Service.FindArticles(id)
		u := user
		u.Articles = articles
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
		category := ""
		url := ""
		if len(r.Form["category"]) > 0 {
			category = r.Form["category"][0]
		}
		if len(r.Form["url"]) > 0 {
			url = r.Form["url"][0]
		}
		id, _ := strconv.Atoi(userID)
		c.Service.SaveArticleToUser(name, author, website, id, category, url)
	})
}

func (c *controller) ShowArticlesAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("homepage.html")
		a := []domain.Article{}
		res, err := http.Get("http://localhost:5050/articles")
		if err != nil {
			panic(err.Error())
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		err = json.Unmarshal(body, &a)
		if err != nil {
			fmt.Println("unmarshall error: ", err)
		}
		user := domain.User{}
		token, err := r.Cookie("hn_auth_token")
		if err == nil {
			user, err = c.Service.FindUserByAuth(token.Value)
		}
		data := struct {
			Articles []domain.Article
			User     domain.User
		}{
			Articles: a,
			User:     user,
		}
		err = t.Execute(w, data)
	})
}
