package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/users")
	routes.GET("/", h.GetAllUsers)
	routes.GET("/:id", h.GetUser)

}
