package articles

import (
	"net/http"
)

// InitializeHandler comment
func InitializeHandler(server *http.ServeMux, as Service) {
	controller := makeController(as)
	server.Handle("/category/", controller.ShowArticlesCategory())
}
