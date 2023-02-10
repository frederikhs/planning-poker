package lobby

import (
	"fmt"
	"github.com/gofiber/websocket/v2"
	"sync"
)

type Client struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Value    string `json:"value"`
	Conn     *websocket.Conn
}

type Lobby struct {
	Name    string
	Clients map[*Client]bool
	lock    *sync.Mutex
}

func (l *Lobby) PrintClients() {
	l.lock.Lock()
	defer l.lock.Unlock()

	fmt.Println("Clients:")
	for client, _ := range l.Clients {
		fmt.Println(client.Id)
	}
}

func (l *Lobby) GetClients() []*Client {
	l.lock.Lock()
	defer l.lock.Unlock()

	var clients []*Client

	for client, _ := range l.Clients {
		clients = append(clients, client)
	}

	return clients
}

func NewLobby(name string) *Lobby {
	return &Lobby{
		Name:    name,
		Clients: make(map[*Client]bool),
		lock:    &sync.Mutex{},
	}
}

func NewClient(conn *websocket.Conn, id string) *Client {
	return &Client{
		Id:       id,
		Username: "username",
		Value:    "5",
		Conn:     conn,
	}
}
