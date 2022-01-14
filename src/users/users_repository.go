package users

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IUserRepository interface {
	saveUser(newUser UserModel) *UserModel
	findAll() *[]UserModel
	findUserById(id uint) *UserModel
	updateUser(newInfo UserModel) *UserModel
	deleteUser(id uint) *UserModel
}

type UserRepository struct {
}

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&UserModel{})
}

func (ur UserRepository) saveUser(newUser UserModel) *UserModel {
	db.Create(&newUser)
	return &newUser
}

func (ur UserRepository) findAll() (allUsers *[]UserModel) {
	db.Find(&allUsers)
	return allUsers
}

func (ur UserRepository) findUserById(id uint) (userFinded *UserModel) {
	db.Find(&userFinded, id)
	return userFinded
}

func (ur UserRepository) updateUser(newInfo UserModel) (userUpdated *UserModel) {
	userFinded := ur.findUserById(newInfo.ID)

	db.Model(&userFinded).Updates(UserModel{Email: newInfo.Email, Password: newInfo.Password})

	return userUpdated
}

func (ur UserRepository) deleteUser(id uint) *UserModel {
	userFinded := ur.findUserById(id)
	db.Delete(&UserModel{}, userFinded.ID)
	return userFinded
}
