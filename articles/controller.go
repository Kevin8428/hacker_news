package articles

import (
	"log"
	"net/http"

	"github.com/alecthomas/template"
)

type ArticlesNewHandler struct{}

func (a ArticlesNewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("homepage.html")
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal("Execute error: ", err)
		return
	}
}
