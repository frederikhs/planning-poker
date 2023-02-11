package hub

import (
	"log"
	"sync"
)

type State struct {
	Sessions map[string]Session
	Hubs     map[string]*Hub
	Lock     *sync.Mutex
}

type Session struct {
	ClientId string `json:"client_id"`
	Username string `json:"username"`
	Value    string `json:"value"`
}

func NewState() *State {
	return &State{
		Sessions: make(map[string]Session),
		Hubs:     make(map[string]*Hub),
		Lock:     &sync.Mutex{},
	}
}

func (s *State) GetOrCreateHub(name string) *Hub {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	if value, exists := s.Hubs[name]; exists {
		return value
	}

	log.Println("creating hub " + name)

	h := NewHub()
	go func() {
		h.Run()

		log.Printf("deleting hub " + name)

		// run blocks, so when it is done delete hub
		s.Lock.Lock()
		delete(s.Hubs, name)
		s.Lock.Unlock()
	}()

	s.Hubs[name] = h

	return h
}

func (s *State) GetSessionClientId(sessionId string) *Session {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	if value, exists := s.Sessions[sessionId]; exists {
		return &value
	}

	return nil
}

func (s *State) SetSession(sessionId, userId string, username string, value string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Sessions[sessionId] = Session{
		ClientId: userId,
		Username: username,
		Value:    value,
	}
}
