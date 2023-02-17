package app

import (
	"encoding/json"
	"github.com/frederikhs/namer"
	"github.com/frederikhs/planning-poker/hub"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func NewSessionCookie(w http.ResponseWriter, s *hub.State) *string {
	sId := uuid.New().String()
	cId := uuid.New().String()

	s.SetSession(sId, cId, namer.GeneratePascalName(sId), -1)

	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: sId,
		Path:  "/",
	})

	return &cId
}

func Create() http.Handler {
	s := hub.NewState()
	r := mux.NewRouter()

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")

		var clientId *string

		if err != nil {
			clientId = NewSessionCookie(w, s)
		} else {
			sessionId := cookie.Value
			session := s.GetSessionClientId(sessionId)

			// if this session id does not match a client, make new session
			if session == nil {
				clientId = NewSessionCookie(w, s)
			} else {
				clientId = &session.ClientId
			}
		}

		b, err := json.Marshal(struct {
			ClientId string `json:"client_id"`
		}{
			ClientId: *clientId,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(b)
	})

	r.HandleFunc("/ws/{lobby_id}", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(s, w, r)
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowCredentials: true,
	})

	return c.Handler(r)
}
