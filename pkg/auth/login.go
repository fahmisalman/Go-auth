package auth

import (
	"go_auth/pkg/common/models"
	"go_auth/pkg/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (h handler) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	if err := h.DB.Model(models.User{}).Where("username = ?", input.Username).Take(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password not valid"})
		return
	}

	if err := verifyPassword(input.Password, user.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password not valid"})
		return
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login!", "access_token": token})
}
