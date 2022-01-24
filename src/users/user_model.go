package users

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	LastLogin      time.Time `json:"last_login"`
	ActivationHash string    `json:"-"`
}

func (user UserModel) string() string {
	return fmt.Sprintf("ID: %v, email: %v, Pass: %v, ActivationHash: %v, LastLogin: %v.",
		user.ID, user.Email, user.Password, user.ActivationHash, user.LastLogin)
}
