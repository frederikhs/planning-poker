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

	valuesVisible bool
}

func NewHub(s *State) *Hub {
	h := &Hub{
		broadcast:     make(chan []byte),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		valuesVisible: false,
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
		ToggleVisibilityRequestEventHandler: func(event ToggleVisibilityRequestEvent) {
			h.valuesVisible = !h.valuesVisible

			b, err := json.Marshal(NewToggleVisibilityEvent(h.valuesVisible))
			if err != nil {
				log.Println(err)
				return
			}

			h.broadcast <- b
		},
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
		ClearLobbyEventHandler: func(event ClearLobbyEvent) {
			s.Lock.Lock()
			for client := range h.clients {
				client.Session.Value = -1
				s.UpdateSessionWithoutLock(client.Session.SessionId, client.Session)
			}
			s.Lock.Unlock()

			h.valuesVisible = false

			b, err := json.Marshal(event)
			if err != nil {
				log.Println(err)
				return
			}

			h.broadcast <- b
		},
		ChooseUsernameHandler: func(client *Client, event ChooseUsernameEvent) {
			client.Session.Username = event.Username

			// store this users' new username
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
