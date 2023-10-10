package user

import (
	"go_auth/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.GetUserByID(id)
	if err != nil {
		h.HandleUserError(c, err)
		return
	}

	h.RespondWithUserJSON(c, user)
}

func (h UserHandler) GetUserByID(id string) (*models.User, error) {
	var user models.User

	result := h.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (h UserHandler) HandleUserError(c *gin.Context, err error) {
	c.AbortWithError(http.StatusNotFound, err)
}

func (h UserHandler) RespondWithUserJSON(c *gin.Context, user *models.User) {
	c.JSON(http.StatusOK, user)
}
