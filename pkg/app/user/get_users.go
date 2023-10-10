package user

import (
	"go_auth/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userResponse struct {
	ID       int    `json:id`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func (h handler) GetAllUsers(c *gin.Context) {
	var users []models.User

	if result := h.DB.Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var userResponseList []userResponse

	for _, users := range users {
		userResponse := userResponse{
			ID:       int(users.ID),
			Username: users.Username,
			Email:    users.Email,
			Name:     users.Name,
		}
		userResponseList = append(userResponseList, userResponse)
	}

	c.JSON(http.StatusOK, userResponseList)
}
