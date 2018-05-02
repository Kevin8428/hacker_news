package main

import (
	"fmt"
	"log"
	"net/http"

	"text/template"
)

func renderPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "homepage.html")
}

type handler struct{}

func count() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

type Person struct {
	Name   string
	Emails []string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := `{{$name := .Name}}
	{{range .Emails}}
			Name is {{$name}}, email is {{.}}
	{{end}}
	`
	person := Person{
		Name:   "Satish",
		Emails: []string{"satish@rubylearning.org", "satishtalim@gmail.com"},
	}
	t := template.New("Person template")
	t, _ = t.Parse(tmpl)
	err := t.Execute(w, person)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}

func buildHandler(page *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page.Execute(w, r)
	})
}

func main() {
	////////////////implement without handler////////////////
	// http.HandleFunc("/", renderPage)
	// err := http.ListenAndServe(":5050", nil)
	// if err != nil {
	// 	panic(err)
	// }

	////////////////implement with handler////////////////
	server := http.NewServeMux()
	server.Handle("/", handler{})
	err := http.ListenAndServe(":5050", server)
	if err != nil {
		panic(err)
	}

	////////////////implement with handler function////////////////
	// server := http.NewServeMux()
	// homepage := template.Must(template.ParseFiles("homepage.html"))
	// server.Handle("/", buildHandler(homepage))
	// if err := http.ListenAndServe(":5050", server); err != nil {
	// 	panic(err)
	// }

}
