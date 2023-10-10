package auth

import (
	"go_auth/pkg/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/auth")
	routes.POST("/register", h.Register)
	routes.POST("/login", h.Login)

	profile := routes.Group("/profile")
	profile.Use(middlewares.JwtAuthMiddleware())
	profile.GET("", h.CurrentUser)
}
