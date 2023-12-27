package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", welcome)
	registerEventRoutes(server)
}

func welcome(context *gin.Context) {
	context.String(http.StatusOK, "Welcome to the server")
}
