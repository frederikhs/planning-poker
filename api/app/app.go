package app

import (
	"github.com/frederikhs/planning-poker/event"
	"github.com/frederikhs/planning-poker/lobby"
	"github.com/frederikhs/planning-poker/state"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"log"
)

func Create() *fiber.App {
	s := state.New()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost",
		AllowCredentials: true,
	}))

	app.Get("/register", func(c *fiber.Ctx) error {
		clientId := s.GetSessionUserId(c.Cookies("session"))
		if clientId == nil {
			sId := uuid.New().String()
			cId := uuid.New().String()
			s.SetSession(sId, cId)
			clientId = &cId
			c.Cookie(&fiber.Cookie{
				Name:  "session",
				Value: sId,
				//HTTPOnly: true,
				//SameSite: http.SameSiteLaxMode,
				Path: "/",
			})
		}

		c.Status(fiber.StatusCreated)

		return c.JSON(struct {
			ClientId string `json:"client_id"`
		}{
			ClientId: *clientId,
		})
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(c) {
			return fiber.ErrUpgradeRequired
		}

		clientId := s.GetSessionUserId(c.Cookies("session"))
		if clientId == nil {
			return fiber.ErrMethodNotAllowed
		}

		return c.Next()
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		v := s.GetSessionUserId(c.Cookies("session"))
		if v == nil {
			c.Close()
			return
		}

		l := s.GetOrCreateLobby("test")

		clientId := *v

		//existingLobby := s.GetClientLobby(clientId)
		//if existingLobby != nil {
		//	existingLobby
		//}

		s.RemoveFromAllLobbies(clientId)

		client := lobby.NewClient(c, clientId, "username", "5")

		_ = c.WriteJSON(event.WelcomeEvent{
			Event:   event.Event{EventType: event.WelcomeEventType},
			Client:  client,
			Clients: l.GetClients(),
		})

		// notify other clients that this one has joined
		_ = l.WriteAll(event.JoinEvent{
			Event:  event.Event{EventType: event.JoinEventType},
			Client: client,
		})

		l.AddClient(client)

		var (
			msg []byte
			err error
		)

		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			log.Printf("recv: %s", msg)
		}

		l.RemoveClient(client)

		_ = l.WriteAll(event.LeaveEvent{
			Event:  event.Event{EventType: event.LeaveEventType},
			Client: client,
		})
	}))

	return app
}
