package users

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type IUserRepository interface {
	saveUser(newUser *UserModel) *UserModel
	findAll() *[]UserModel
	findUserById(id uint) *UserModel
	updateUser(oldUser *UserModel, newInfo *UserModel) *UserModel
	deleteUser(id uint) *UserModel
	findByEmail(email string) *UserModel
}

type UserRepository struct {
}

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&UserModel{})
}

func (ur UserRepository) saveUser(newUser *UserModel) *UserModel {
	db.Create(&newUser)
	return newUser
}

func (ur UserRepository) findAll() (allUsers *[]UserModel) {
	db.Last(&allUsers)
	return allUsers
}

func (ur UserRepository) findUserById(id uint) (userFinded *UserModel) {
	db.Last(&userFinded, id)
	return userFinded
}

func (ur UserRepository) findByEmail(email string) (userFinded *UserModel) {
	db.Where("Email = ?", email).Last(&userFinded)
	if userFinded.ID == 0 || userFinded.Email == "" {
		return nil
	}
	return userFinded
}

func (ur UserRepository) updateUser(oldUser *UserModel, newInfo *UserModel) (userUpdated *UserModel) {
	db.Model(&oldUser).Updates(newInfo)
	return oldUser
}

func (ur UserRepository) deleteUser(id uint) *UserModel {
	userFinded := ur.findUserById(id)
	db.Delete(&UserModel{}, userFinded.ID)
	return userFinded
}
