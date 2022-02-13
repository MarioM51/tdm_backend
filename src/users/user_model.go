package users

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email          string      `json:"email"`
	Password       string      `json:"password,omitempty"`
	LastLogin      time.Time   `json:"last_login,omitempty"`
	ActivationHash string      `json:"-"`
	Rols           []RoleModel `gorm:"many2many:user_roles;" json:"rols"`
}

type RoleModel struct {
	gorm.Model
	Name string `gorm:"unique"`
}

func (user UserModel) string() string {
	return fmt.Sprintf("ID: %v, email: %v, Pass: %v, ActivationHash: %v, LastLogin: %v.",
		user.ID, user.Email, user.Password, user.ActivationHash, user.LastLogin)
}
