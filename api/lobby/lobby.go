package lobby

import (
	"fmt"
	"log"
	"sync"
)

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

	clients := []*Client{}

	for client, _ := range l.Clients {
		clients = append(clients, client)
	}

	return clients
}

func (l *Lobby) AddClient(c *Client) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.Clients[c] = true
}

func (l *Lobby) GetClientById(id string) *Client {
	l.lock.Lock()
	defer l.lock.Unlock()

	for client, _ := range l.Clients {
		if client.Id == id {
			return client
		}
	}

	return nil
}

func (l *Lobby) RemoveClient(c *Client) {
	l.lock.Lock()
	defer l.lock.Unlock()

	log.Printf("removing client " + c.Id)

	err := c.Conn.Close()
	if err != nil {
		log.Println(err)
	}
	delete(l.Clients, c)
}

func (l *Lobby) WriteAll(i interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	for client, _ := range l.Clients {
		err := client.Conn.WriteJSON(i)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewLobby(name string) *Lobby {
	return &Lobby{
		Name:    name,
		Clients: make(map[*Client]bool),
		lock:    &sync.Mutex{},
	}
}
