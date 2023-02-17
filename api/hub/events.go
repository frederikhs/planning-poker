package hub

import (
	"encoding/json"
	"errors"
)

const (
	JoinEventType                    = "join_event"
	LeaveEventType                   = "leave_event"
	ChooseUsernameEventType          = "choose_username_event"
	PickEventType                    = "pick_event"
	SessionChangeEventType           = "session_change_event"
	WelcomeEventType                 = "welcome_event"
	ClearLobbyEventType              = "clear_lobby_event"
	ToggleVisibilityRequestEventType = "toggle_visibility_request_event"
	ToggleVisibilityEventType        = "toggle_visibility_event"
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

type SessionChangeEvent struct {
	Event
	Session *Session `json:"session"`
}

func NewSessionChangeEvent(session *Session) SessionChangeEvent {
	return SessionChangeEvent{
		Event:   Event{EventType: SessionChangeEventType},
		Session: session,
	}
}

type PickEvent struct {
	Event
	Value int `json:"value"`
}

type WelcomeEvent struct {
	Event
	Session  *Session   `json:"session"`
	Sessions []*Session `json:"sessions"`
	Visible  bool       `json:"visible"`
}

func NewWelcomeEvent(session *Session, sessions []*Session, visible bool) WelcomeEvent {
	return WelcomeEvent{
		Event:    Event{EventType: WelcomeEventType},
		Session:  session,
		Sessions: sessions,
		Visible:  visible,
	}
}

func NewWelcomeEventFromHub(hub *Hub, session *Session) WelcomeEvent {
	sessions := []*Session{}
	for someClient, _ := range hub.clients {
		sessions = append(sessions, someClient.Session)
	}

	return NewWelcomeEvent(session, sessions, hub.valuesVisible)
}

type ToggleVisibilityEvent struct {
	Event
	Visible bool `json:"visible"`
}

func NewToggleVisibilityEvent(visible bool) ToggleVisibilityEvent {
	return ToggleVisibilityEvent{
		Event:   Event{EventType: ToggleVisibilityEventType},
		Visible: visible,
	}
}

type ToggleVisibilityRequestEvent struct {
	Event
}

type ClearLobbyEvent = Event

type EventHandler struct {
	ChooseUsernameHandler               func(event ChooseUsernameEvent)
	JoinEventHandler                    func(event JoinEvent)
	LeaveEventHandler                   func(event LeaveEvent)
	PickEventHandler                    func(client *Client, event PickEvent)
	ClearLobbyEventHandler              func(event ClearLobbyEvent)
	ToggleVisibilityRequestEventHandler func(event ToggleVisibilityRequestEvent)
}

func (eh EventHandler) Handle(client *Client, message []byte) error {
	var e Event
	err := json.Unmarshal(message, &e)
	if err != nil {
		return err
	}

	switch e.EventType {
	//case JoinEventType:
	//	var e JoinEvent
	//	err := json.Unmarshal(message, &e)
	//	if err != nil {
	//		return err
	//	}
	//
	//	eh.JoinEventHandler(e)
	//	return nil
	//case LeaveEventType:
	//	var e LeaveEvent
	//	err := json.Unmarshal(message, &e)
	//	if err != nil {
	//		return err
	//	}
	//
	//	eh.LeaveEventHandler(e)
	//	return nil
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

		eh.PickEventHandler(client, e)
		return nil
	//case ClearLobbyEventType:
	//	var e ClearLobbyEvent
	//	err := json.Unmarshal(message, &e)
	//	if err != nil {
	//		return err
	//	}
	//
	//	eh.ClearLobbyEventHandler(e)
	//	return nil
	case ToggleVisibilityRequestEventType:
		var e ToggleVisibilityRequestEvent
		err := json.Unmarshal(message, &e)
		if err != nil {
			return err
		}

		eh.ToggleVisibilityRequestEventHandler(e)
		return nil
	default:
		return errors.New("cannot handle event type:" + e.EventType)
	}
}
