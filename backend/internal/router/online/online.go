package online

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func Online(c echo.Context) error {
	//tokenString := c.QueryParam("token")
	//fmt.Println(tokenString)
	//userID, err := token.GetUserIDFromTokenString(tokenString)
	//fmt.Println(userID, err)
	//if err != nil {
	//	return c.String(http.StatusUnauthorized, err.Error())
	//}

	conn, _ := upgrader.Upgrade(c.Response(), c.Request(), nil)

	defer conn.Close()

	//database, closeDB := db.Init()
	//defer closeDB()

	for {

		//database.Model(&model.User{}).Where("id = ?", userID).Update("online", true)

		mt, _, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			//database.Model(&model.User{}).Where("id = ?", userID).Update("online", false)
			break
		}
	}
	return nil
}
