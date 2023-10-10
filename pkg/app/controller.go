package app

import (
	"go_auth/pkg/app/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	routes := r.Group("/app")
	user.RegisterRoutes(routes, db)

}
