package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
var a = 0
var b = 0

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(messageType);
		m := make(map[string]int)
		err = json.Unmarshal(p, &m)
		if val, ok := m["a"]; ok {
			//do something here
			a = val
		}
		if val, ok := m["b"]; ok {
			//do something here
			b = val
		}
		if val, ok := m["sum"]; ok {
			log.Println(val)
			log.Println("Sending computed sum back")
			values := map[string]int{"sum": (a+b)}
			jsonData, err := json.Marshal(values)
			log.Println(jsonData)
			if err == nil {
				conn.WriteMessage(1,jsonData)
			}
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
