package websockets

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type connection struct {
	// Buffered channel of outbound messages.
	send chan []byte

	// The hub.
	h *hub
}

func (c *connection) writer(wg *sync.WaitGroup, wsConn *websocket.Conn) {
	defer wg.Done()
	for message := range c.send { // listening to send channel
		fmt.Println("one message being sent")
		err := wsConn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
}

func (c *connection) reader(wg *sync.WaitGroup, wsConn *websocket.Conn, isAdmin bool) {
	defer wg.Done()
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("reading!")
		if isAdmin {
			c.h.adminChan <- message
		} else {
			c.h.broadcast <- message
		}
	}
}

func Connect(hub *hub) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("attempting connection")
		wsConn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("error upgrading %s", err)
			return
		}
		c := &connection{send: make(chan []byte, 256), h: hub}
		c.h.addHostConnection(c)
		defer c.h.removeHostConnection(c)
		var wg sync.WaitGroup
		wg.Add(2)
		go c.writer(&wg, wsConn)
		go c.reader(&wg, wsConn, true)
		wg.Wait()
		wsConn.Close()
	})
}
