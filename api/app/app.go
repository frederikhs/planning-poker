package app

import (
	"encoding/json"
	"fmt"
	"github.com/frederikhs/namer"
	"github.com/frederikhs/planning-poker/hub"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"html"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// stolen
func randomEmoji() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	// http://apps.timwhitlock.info/emoji/tables/unicode
	emoji := [][]int{
		// Emoticons icons
		{128513, 128591},
		// Transport and map symbols
		{128640, 128704},
	}
	r := emoji[rand.Int()%len(emoji)]
	min := r[0]
	max := r[1]
	n := rand.Intn(max-min+1) + min
	return html.UnescapeString("&#" + strconv.Itoa(n) + ";")
}

func NewSessionCookie(w http.ResponseWriter, s *hub.State) *string {
	sId := uuid.New().String()
	cId := uuid.New().String()

	s.SetSession(sId, cId, randomEmoji()+" "+namer.GeneratePascalName(sId), -1, false)

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sId,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   60 * 60 * 24 * 365, // 1 year
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
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	})

	r.HandleFunc("/lobby/create", func(w http.ResponseWriter, r *http.Request) {
		name := uuid.New().String()
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(fmt.Sprintf("{\"lobby_id\": \"%s\"}", name)))
		if err != nil {
			log.Println(err)
		}
	})

	r.HandleFunc("/ws/{lobby_id}", func(w http.ResponseWriter, r *http.Request) {
		hub.ServeWs(s, w, r)
	})

	r.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(fmt.Sprintf("{\"version\": \"%s\"}", os.Getenv("VERSION"))))
		if err != nil {
			log.Println(err)
		}
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "https://pp.hrgn.dk"},
		AllowCredentials: true,
	})

	return c.Handler(r)
}
