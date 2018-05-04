package articles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"text/template"

	"github.com/kevin8428/hackernews/domain"
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
		err = t.Execute(w, a)
	})
}
