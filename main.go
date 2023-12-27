package main

import (
	"github.com/duyanhitbe/go-basic-rest/database"
	"github.com/duyanhitbe/go-basic-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
