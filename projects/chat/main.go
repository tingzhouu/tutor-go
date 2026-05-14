package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	connections map[*websocket.Conn]string = make(map[*websocket.Conn]string)
	mu          sync.Mutex
)

func addConnection(c *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	connections[c] = ""
	fmt.Printf("we have %d connections\n", len(connections))
	fmt.Println(connections)
}

func broadcast(message string) {
	for connection := range connections {
		if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			fmt.Println(err)
			delete(connections, connection)
			connection.Close()
		}
	}
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
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			mu.Lock()
			broadcast(connections[conn] + " has left")
			delete(connections, conn)
			mu.Unlock()
			conn.Close()
			return
		}

		mu.Lock()
		if connections[conn] == "" {
			connections[conn] = string(p)
			broadcast(string(p) + " has joined")
		} else {
			broadcast(fmt.Sprintf("[%s]: %s", connections[conn], string(p)))
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
