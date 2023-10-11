package auth

import (
	"fmt"
	"go_auth/pkg/app/user"
	"go_auth/pkg/common/models"
	"go_auth/pkg/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get the current user
// @Description Get information about the currently authenticated user.
// @Produce json
// @Success 200 {object} CurrentUserResponse
// @Failure 400 {object} ErrorResponse
// @Router /auth/profile [get]
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

// CurrentUserResponse represents the response structure for the CurrentUser endpoint.
type CurrentUserResponse struct {
	Message string          `json:"message"`
	Data    CurrentUserData `json:"data"`
}

// CurrentUserData represents the information about the current user.
type CurrentUserData struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h handler) getCurrentUser(userID uint) (*models.User, error) {
	uh := &user.UserHandler{
		DB: h.DB,
	}

	user, err := uh.GetUserByID(fmt.Sprintf("%d", userID))
	return user, err
}
