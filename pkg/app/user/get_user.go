package user

import (
	"go_auth/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &user)
}
