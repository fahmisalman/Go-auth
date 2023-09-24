package main

import (
	auth "go_auth/auth"
	config "go_auth/config"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	router := gin.Default()

	auth.RegisterRoutes(router, db)

	router.Run()
}
