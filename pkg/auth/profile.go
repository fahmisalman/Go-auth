package auth

import (
	"fmt"
	"go_auth/pkg/app/user"
	"go_auth/pkg/common/models"
	"go_auth/pkg/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) CurrentUser(c *gin.Context) {
	userID, tokenErr := token.ExtractTokenID(c)
	if tokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tokenErr.Error()})
		return
	}

	currentUser, userErr := h.getCurrentUser(userID)
	if userErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": userErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": currentUser})
}

func (h handler) getCurrentUser(userID uint) (*models.User, error) {
	uh := &user.UserHandler{
		DB: h.DB,
	}

	user, err := uh.GetUserByID(fmt.Sprintf("%d", userID))
	return user, err
}
