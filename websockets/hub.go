package websockets

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

// Constructor
func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

func (h *Hub) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// utils.ErrorResponse(500, "Could not open the websocket connection", w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	client := NewClient(h, socket)

	//Client register
	h.register <- client

	go client.Write()

}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}

func (hub *Hub) onConnect(client *Client) {
	log.Println("A new client connected: ", client.socket.RemoteAddr())

	//To avoid race condition
	client.hub.mutex.Lock()

	defer client.hub.mutex.Unlock()

	client.id = client.socket.RemoteAddr().String()

	//Add the client to clients slice
	hub.clients = append(hub.clients, client)

}

func (hub *Hub) onDisconnect(client *Client) {
	log.Println("Client disconnected: ", client.socket.RemoteAddr())

	client.Close()
	client.hub.mutex.Lock()
	defer client.hub.mutex.Unlock()

	i := -1

	for j, c := range hub.clients {
		if c.id == client.id {
			i = j //Copy the  index
			break
		}
	}

	copy(hub.clients[i:], hub.clients[i+1:])
	hub.clients[len(hub.clients)-1] = nil
	hub.clients = hub.clients[:len(hub.clients)-1]
}

func (hub *Hub) BroadCast(message any, ignore *Client) {
	data, _ := json.Marshal(message)

	for _, client := range hub.clients {
		if client != ignore {
			client.outbound <- data
		}
	}
}
