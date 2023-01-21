package main

import (
	"real-time-chat/configuration"
	"real-time-chat/routes"
	"real-time-chat/websocketComponent"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configuration.ConnectDB()

	routes.UserRoute(router)

	go websocketComponent.Hubs.Run()

	router.Run()
}
