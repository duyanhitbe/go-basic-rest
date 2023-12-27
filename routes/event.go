package routes

import (
	"github.com/duyanhitbe/go-basic-rest/controllers"
	"github.com/gin-gonic/gin"
)

func registerEventRoutes(server *gin.Engine) {
	server.POST("/events", controllers.CreateEvent)
	server.GET("/events", controllers.GetAllEvent)
	server.GET("/events/:id", controllers.GetOneEvent)
	server.PUT("/events/:id", controllers.UpdateOneEvent)
	server.DELETE("/events/:id", controllers.DeleteOneEventById)
}
