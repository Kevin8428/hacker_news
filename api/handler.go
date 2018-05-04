package api

import "net/http"

// InitializeHandler comment
func InitializeHandler(server *http.ServeMux) {
	server.Handle("/articles", ShowAllArticles())
}
