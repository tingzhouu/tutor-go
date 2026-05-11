package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	connections map[*websocket.Conn]bool = make(map[*websocket.Conn]bool)
	mu          sync.Mutex
)

func addConnection(c *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	connections[c] = true
	fmt.Printf("we have %d connections\n", len(connections))
	fmt.Println(connections)
}

func startWebsocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	addConnection(conn)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			mu.Lock()
			delete(connections, conn)
			mu.Unlock()
			conn.Close()
			return
		}

		mu.Lock()
		for connection := range connections {
			if err := connection.WriteMessage(messageType, p); err != nil {
				fmt.Println(err)
				delete(connections, connection)
				connection.Close()
			}
		}
		mu.Unlock()

	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {

	http.HandleFunc("GET /ws", startWebsocket)
	http.HandleFunc("GET /", homePage)

	http.ListenAndServe(":3000", nil)
}
