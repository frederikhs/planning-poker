package state

import (
	"github.com/frederikhs/planning-poker/lobby"
	"log"
	"sync"
)

type State struct {
	Sessions map[string]string
	Lobbies  map[string]*lobby.Lobby
	lock     *sync.Mutex
}

func New() *State {
	return &State{
		Sessions: make(map[string]string),
		Lobbies:  make(map[string]*lobby.Lobby),
		lock:     &sync.Mutex{},
	}
}

func (s *State) GetSessionUserId(sessionId string) *string {
	s.lock.Lock()
	defer s.lock.Unlock()

	if value, exists := s.Sessions[sessionId]; exists {
		return &value
	}

	return nil
}

func (s *State) SetSession(sessionId, userId string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Sessions[sessionId] = userId
}

func (s *State) GetOrCreateLobby(id string) *lobby.Lobby {
	s.lock.Lock()
	defer s.lock.Unlock()

	if value, exists := s.Lobbies[id]; exists {
		return value
	}

	l := lobby.NewLobby(id)
	s.Lobbies[id] = l

	return l
}

func (s *State) RemoveFromAllLobbies(clientId string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, l := range s.Lobbies {
		for client, _ := range l.Clients {
			if client.Id == clientId {
				log.Println("removing client from lobby")
				l.RemoveClient(client)
			}
		}
	}
}
