package main

import (
	config "go_auth/config"
	auth "go_auth/module/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	router := gin.Default()

	auth.RegisterRoutes(router, db)

	router.Run()
}
