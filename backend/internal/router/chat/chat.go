package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/token"
)

var (
	upgrader = websocket.Upgrader{}
)

var Clients = make(map[uint]*websocket.Conn)

func Chat(c echo.Context) error {
	userID := token.GetUserIDFromToken(&c)
	conn, _ := upgrader.Upgrade(c.Response(), c.Request(), nil)

	defer conn.Close()

	Clients[userID] = conn
	fmt.Println(Clients)

	for {

		// Read
		mt, _, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			delete(Clients, userID)
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}
	}
	return nil
}
