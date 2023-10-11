package main

import (
	_ "go_auth/docs"
	app "go_auth/pkg/app"
	auth "go_auth/pkg/auth"
	db "go_auth/pkg/common/db"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()
	router := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	auth.RegisterRoutes(router, database)
	app.RegisterRoutes(router, database)

	router.Run()
}
