package app

import (
	"encoding/json"
	"github.com/frederikhs/planning-poker/lobby"
	"github.com/frederikhs/planning-poker/state"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"log"
)

func Create() *fiber.App {
	s := state.New()

	app := fiber.New()

	l := lobby.NewLobby("test")

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

		clientId := *v

		client := lobby.NewClient(c, clientId)
		l.Clients[client] = true

		l.PrintClients()

		b, err := json.Marshal(struct {
			ClientId string          `json:"client_id"`
			Clients  []*lobby.Client `json:"clients"`
		}{
			ClientId: clientId,
			Clients:  l.GetClients(),
		})
		if err != nil {
			return
		}

		err = c.WriteMessage(1, b)
		if err != nil {
			return
		}

		var (
			mt  int
			msg []byte
			//err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}

			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}

		delete(l.Clients, client)

	}))

	return app
}
