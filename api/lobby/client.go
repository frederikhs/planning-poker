package lobby

import "github.com/gofiber/websocket/v2"

type Client struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Value    string `json:"value"`
	Conn     *websocket.Conn
}

func NewClient(conn *websocket.Conn, id, username, value string) *Client {
	return &Client{
		Id:       id,
		Username: username,
		Value:    value,
		Conn:     conn,
	}
}
