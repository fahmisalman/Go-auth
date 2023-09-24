package main

import (
	controllers "go_auth/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	auth := r.Group("/auth")
	auth.POST("/register", controllers.Register)

	r.Run()
}
