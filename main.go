package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satyam-jha-16/jwt-auth/controllers"
	"github.com/satyam-jha-16/jwt-auth/initializers"
	"github.com/satyam-jha-16/jwt-auth/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health": "ok",
		})
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
