package event

import (
	"encoding/json"
	"errors"
	"github.com/frederikhs/planning-poker/lobby"
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
	Client *lobby.Client `json:"client"`
}

type LeaveEvent struct {
	Event
	Client *lobby.Client `json:"client"`
}

type ChooseUsernameEvent struct {
	Event
	Username string `json:"username"`
}

type UserChangeEvent struct {
	Event
	Client *lobby.Client `json:"client"`
}

type PickEvent struct {
	Event
	Value int `json:"value"`
}

type WelcomeEvent struct {
	Event
	Client  *lobby.Client   `json:"client"`
	Clients []*lobby.Client `json:"clients"`
}

type ToggleVisibilityEvent struct {
	Event
	Visible bool `json:"visible"`
}

type ClearLobbyEvent = Event

type Handler struct {
	ChooseUsernameHandler  func(event ChooseUsernameEvent)
	JoinEventHandler       func(event JoinEvent)
	LeaveEventHandler      func(event LeaveEvent)
	PickEventHandler       func(event PickEvent)
	ClearLobbyEventHandler func(event ClearLobbyEvent)
	VisibilityEventHandler func(event ToggleVisibilityEvent)
}

func (eh Handler) Handle(message []byte) error {
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
