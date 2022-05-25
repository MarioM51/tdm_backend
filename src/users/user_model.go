package users

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email          string `json:"email"`
	Password       string `json:"password,omitempty"`
	ActivationHash string `json:"-"`
	//Rols           []RoleModel `gorm:"many2many:users_rols;" json:"rols"`
	Rols []RoleModel `json:"rols" gorm:"many2many:users_rols;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:role_id"`
}

// Which creates join table: users_rols
//   foreign key: user_id, reference: users.id
//   foreign key: profile_refer, reference: profiles.user_refer

func (UserModel) TableName() string {
	return "users"
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
