package main

import (
	app "go_auth/pkg/app"
	auth "go_auth/pkg/auth"
	db "go_auth/pkg/common/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()
	router := gin.Default()

	auth.RegisterRoutes(router, database)
	app.RegisterRoutes(router, database)

	router.Run()
}
