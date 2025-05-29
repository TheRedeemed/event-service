package main

import (
	"the-redeemed/event-service/db"
	"the-redeemed/event-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() //configures an http server

	routes.RegisterRoutes(server)

	server.Run(":8080") //configure server port
}
