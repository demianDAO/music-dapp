package main

import (
	"github.com/gin-gonic/gin"
	"web3-music-platform/internal/routes"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
