package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type Server struct {
	clients       map[*websocket.Conn]bool
	handleMessage func(message []byte) // хандлер новых сообщений
}

func Hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))

		// Read
		mt, msg, err := ws.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}
		fmt.Printf("%s\n", msg)
	}
	return nil
}
