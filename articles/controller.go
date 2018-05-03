package articles

import (
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

func (c *controller) ShowArticles() http.Handler {
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
