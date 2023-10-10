package app

import (
	"go_auth/pkg/app/user"
	"go_auth/pkg/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	routes := r.Group("/app")
	routes.Use(middlewares.JwtAuthMiddleware())

	user.RegisterRoutes(routes, db)
}
