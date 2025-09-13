package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Initialize Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "order-service",
		})
	})

	// Basic order endpoints (to be implemented)
	r.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get orders endpoint - to be implemented",
		})
	})

	r.POST("/orders", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Create order endpoint - to be implemented",
		})
	})

	// Start server
	log.Printf("Order service starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
