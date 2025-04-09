package main

import (
	"myapp/controllers"
	"myapp/database"
	"myapp/middleware"

	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.InitDB()

	r := gin.Default()

	// Public Routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Protected Routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/users", controllers.GetUsers)
	protected.GET("/users/:id", controllers.GetUserByID)
	protected.PUT("/users/:id", controllers.UpdateUser)
	protected.DELETE("/users/:id", controllers.DeleteUser)

	log.Println("🚀 Server is running on port 8080")
	r.Run(":8080")
}