package users

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Last_login time.Time `json:"last_login"`
}

func (user UserModel) string() string {
	return fmt.Sprintf("ID: %v, email: %v, Pass: %v", user.ID, user.Email, user.Password)
}
