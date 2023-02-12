package hub

import (
	"encoding/json"
	"log"
	"strconv"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	eventHandler *EventHandler
}

func NewHub(s *State) *Hub {
	h := &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}

	h.eventHandler = CreateEventHandler(h, s)

	return h
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func CreateEventHandler(h *Hub, s *State) *EventHandler {
	return &EventHandler{
		//ChooseUsernameHandler:  nil,
		PickEventHandler: func(client *Client, event PickEvent) {
			if !IsValidValue(event.Value) {
				log.Println(client.Session.ClientId + " picked illegal value " + strconv.Itoa(event.Value))
				return
			}

			log.Println(client.Session.ClientId + " picked " + strconv.Itoa(event.Value))

			client.Session.Value = event.Value

			// store this users choice
			s.UpdateSession(client.Session.SessionId, client.Session)

			b, err := json.Marshal(NewSessionChangeEvent(client.Session))
			if err != nil {
				log.Println(err)
				return
			}

			h.broadcast <- b
		},
	}
}
