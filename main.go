package main

import (
	"example.com/go-jwt/controllers"
	"example.com/go-jwt/initializes"
	"example.com/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializes.LocalEnvVariable()
	initializes.ConnectToDataBase()
	initializes.SyncDataBase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
