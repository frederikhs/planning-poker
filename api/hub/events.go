package hub

import (
	"encoding/json"
	"errors"
)

const (
	JoinEventType             = "join_event"
	LeaveEventType            = "leave_event"
	ChooseUsernameEventType   = "choose_username_event"
	PickEventType             = "pick_event"
	UserChangeEventType       = "user_change_event"
	WelcomeEventType          = "welcome_event"
	ClearLobbyEventType       = "clear_lobby_event"
	ToggleVisibilityEventType = "toggle_visibility_event"
)

type Event struct {
	EventType string `json:"event_type"`
}

type JoinEvent struct {
	Event
	Session *Session `json:"session"`
}

func NewJoinEvent(session *Session) JoinEvent {
	return JoinEvent{
		Event:   Event{EventType: JoinEventType},
		Session: session,
	}
}

type LeaveEvent struct {
	Event
	Session *Session `json:"session"`
}

func NewLeaveEvent(session *Session) LeaveEvent {
	return LeaveEvent{
		Event:   Event{EventType: LeaveEventType},
		Session: session,
	}
}

type ChooseUsernameEvent struct {
	Event
	Username string `json:"Username"`
}

type UserChangeEvent struct {
	Event
	Client string `json:"client"`
}

type PickEvent struct {
	Event
	Value int `json:"value"`
}

type WelcomeEvent struct {
	Event
	Session  *Session   `json:"session"`
	Sessions []*Session `json:"sessions"`
}

func NewWelcomeEvent(session *Session, sessions []*Session) WelcomeEvent {
	return WelcomeEvent{
		Event:    Event{EventType: WelcomeEventType},
		Session:  session,
		Sessions: sessions,
	}
}

func NewWelcomeEventFromHub(hub *Hub, session *Session) WelcomeEvent {
	sessions := []*Session{}
	for someClient, _ := range hub.clients {
		sessions = append(sessions, someClient.Session)
	}

	return NewWelcomeEvent(session, sessions)
}

type ToggleVisibilityEvent struct {
	Event
	Visible bool `json:"visible"`
}

type ClearLobbyEvent = Event

type EventHandler struct {
	ChooseUsernameHandler  func(event ChooseUsernameEvent)
	JoinEventHandler       func(event JoinEvent)
	LeaveEventHandler      func(event LeaveEvent)
	PickEventHandler       func(event PickEvent)
	ClearLobbyEventHandler func(event ClearLobbyEvent)
	VisibilityEventHandler func(event ToggleVisibilityEvent)
}

func (eh EventHandler) Handle(message []byte) error {
	var e Event
	err := json.Unmarshal(message, &e)
	if err != nil {
		return err
	}

	switch e.EventType {
	case JoinEventType:
		var e JoinEvent
		err := json.Unmarshal(message, &e)
		if err != nil {
			return err
		}

		eh.JoinEventHandler(e)
		return nil
	case LeaveEventType:
		var e LeaveEvent
		err := json.Unmarshal(message, &e)
		if err != nil {
			return err
		}

		eh.LeaveEventHandler(e)
		return nil
	case ChooseUsernameEventType:
		var e ChooseUsernameEvent
		err := json.Unmarshal(message, &e)
		if err != nil {
			return err
		}

		eh.ChooseUsernameHandler(e)
		return nil
	case PickEventType:
		var e PickEvent
		err := json.Unmarshal(message, &e)
		if err != nil {
			return err
		}

		eh.PickEventHandler(e)
		return nil
	case ClearLobbyEventType:
		var e ClearLobbyEvent
		err := json.Unmarshal(message, &e)
		if err != nil {
			return err
		}

		eh.ClearLobbyEventHandler(e)
		return nil
	case ToggleVisibilityEventType:
		var e ToggleVisibilityEvent
		err := json.Unmarshal(message, &e)
		if err != nil {
			return err
		}

		eh.VisibilityEventHandler(e)
		return nil
	default:
		return errors.New("cannot handle event type:" + e.EventType)
	}
}
