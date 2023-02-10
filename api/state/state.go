package state

import "sync"

type State struct {
	Sessions map[string]string
	lock     *sync.Mutex
}

func New() *State {
	return &State{
		Sessions: make(map[string]string),
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
