package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Server running!"})
}
