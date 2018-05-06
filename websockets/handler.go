package websockets

import "net/http"

func InitializeHandler(server *http.ServeMux, h *hub) {
	server.Handle("/homepage-ws", Connect(h))
}
