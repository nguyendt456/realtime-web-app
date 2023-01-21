package controllers

import (
	"fmt"
	"net/http"
	"real-time-chat/response"

	"real-time-chat/websocketComponent"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ws = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketController() gin.HandlerFunc {
	return func(c *gin.Context) {
		func(w http.ResponseWriter, r *http.Request) {
			conn, err := ws.Upgrade(w, r, nil)
			if err != nil {
				c.JSON(http.StatusUpgradeRequired, response.UserResponse{Status: http.StatusUpgradeRequired, Message: "Error upgrade websocket", Data: map[string]interface{}{"data": err.Error()}})
				fmt.Println(err.Error())
				return
			}

			client := websocketComponent.ConstructClient(&websocketComponent.Hubs, conn, make(chan []byte))
			websocketComponent.Hubs.Registry(client)

			go client.Read()
			go client.Write()
		}(c.Writer, c.Request)
	}
}
