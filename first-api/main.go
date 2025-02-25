/*
=============================================================================
DEVELOPER: Aswin KS
DATE: 24-02-2025

ABOUT: My first go API
PURPOSE: API that accepts user data and print it
============================================================================
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	//--------------------------------------------------------------------------

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User created Successfully", "user": user,
		})
	})

	//--------------------------------------------------------------------------

	// Start the server on port 8080
	r.Run(":8080")
}
