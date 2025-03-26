package main

import (
	// "log"
	// "virtual-campus-tour-2.0-back/internal/handler"
	// "virtual-campus-tour-2.0-back/internal/middleware"
	// "virtual-campus-tour-2.0-back/pkg/database"
	// "virtual-campus-tour-2.0-back/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize logger
	// logger.Init()

	// Initialize database
	// if err := database.Init(); err != nil {
	// 	log.Fatalf("Failed to initialize database: %v", err)
	// }

	// Create Gin engine
	r := gin.Default()

	// Add middleware
	// r.Use(middleware.Cors())
	// r.Use(middleware.Logger())
	// r.Use(middleware.Recovery())

	// Initialize routes
	// handler.InitRoutes(r)

	// Start server
	r.Run(":8080")
}
