package routes

import (
	"net/http"
	"real-time-chat/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.LoadHTMLFiles("index.html")
	router.POST("/user", controllers.AddUser())
	router.GET("/getuser", controllers.GetUser())
	router.GET("/ws", controllers.WebSocketController())
	router.GET("/chat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
