package articles

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"text/template"
)

type controller struct {
	Service
}

func makeController(as Service) controller {
	return controller{
		Service: as,
	}
}

func (c *controller) ShowArticlesCategory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// these get you to the same place:
		// c.GetArticleInfo()
		// c.Service.GetArticleInfo()
		category := strings.TrimPrefix(r.URL.Path, "/category/")
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		t, err := template.ParseFiles(wd + "/category.html")
		if err != nil {
			log.Fatal("error: ", err)
			return
		}
		err = t.Execute(w, category)
		if err != nil {
			log.Fatal("error: ", err)
			return
		}
	})
}

func (c *controller) ShowSportsCategory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("sports_page.html")
		articles := c.Service.FindArticles()
		fmt.Println("articles: ", articles)
		err = t.Execute(w, articles)
		if err != nil {
			log.Fatal("error: ", err)
			return
		}
	})
}
