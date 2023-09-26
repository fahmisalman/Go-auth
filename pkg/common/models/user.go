package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Name     string `gorm:"size:255;not null;" json:"name"`
	Password string `gorm:"size:255;not null;" json:"password"`
}
