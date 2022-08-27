package users

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email          string      `json:"email"`
	Password       string      `json:"password,omitempty"`
	ActivationHash string      `json:"-"`
	Rols           []RoleModel `json:"rols" gorm:"many2many:users_rols;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:role_id"`
	FullName       string      `json:"fullName"`
	Phone          string      `json:"phone"`
	Zip            string      `json:"zip"`
	State          string      `json:"state"`
	City           string      `json:"city"`
	Street         string      `json:"street"`
	StreetNum      string      `json:"streetNum"`
}

func (UserModel) TableName() string {
	return "users"
}

func (u UserModel) CanConfirmOrder() bool {
	var can = false

	if u.FullName != "" && u.Phone != "" {
		can = true
	}

	return can
}

type RoleModel struct {
	gorm.Model
	Name string `gorm:"unique"`
}

func (RoleModel) TableName() string {
	return "rols"
}

/*
func (user UserModel) String() string {
	return fmt.Sprintf("ID: %v, email: %v, Pass: %v, ActivationHash: %v, LastLogin: %v.",
		user.ID, user.Email, user.Password, user.ActivationHash, user.LastLogin)
}
*/
