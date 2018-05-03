package articles

import "net/http"

// InitializeHandler comment
func InitializeHandler(server *http.ServeMux, as Service) {
	server.Handle("articles-new", ArticlesNewHandler{})
}
