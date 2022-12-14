package users

import (
	"users_api/src/errorss"

	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	saveUser(newUser *UserModel) *UserModel
	findAll() *[]UserModel
	findUserById(id uint) *UserModel
	updateUser(oldUser *UserModel, newInfo *UserModel) *UserModel
	deleteUser(id uint)
	findByEmail(email string) *UserModel
}

type UserRepository struct {
}

func (ur UserRepository) saveUser(newUser *UserModel) *UserModel {
	newUser.Rols = nil
	tx := dbHelper.DB.Create(&newUser)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error saving user"})
	}
	return newUser
}

func (ur UserRepository) findAll() (allUsers *[]UserModel) {
	dbHelper.DB.Preload(clause.Associations).Find(&allUsers)
	for i := range *allUsers {
		(*allUsers)[i].Password = ""
	}
	return allUsers
}

func (UserRepository) findUserById(id uint) *UserModel {
	var userFinded *UserModel
	dbHelper.DB.Preload(clause.Associations).Find(&userFinded, id)
	if userFinded.ID <= 0 {
		return nil
	}
	return userFinded
}

func (ur UserRepository) findByEmail(email string) (userFinded *UserModel) {
	dbHelper.DB.Preload(clause.Associations).Where("Email = ?", email).Last(&userFinded)
	if userFinded.ID == 0 || userFinded.Email == "" {
		return nil
	}
	return userFinded
}

func (ur UserRepository) updateUser(oldUser *UserModel, newInfo *UserModel) (userUpdated *UserModel) {

	newInfo.ID = 0

	//TODO: Agregar a notas, el como hacer actualizacion de una relacion muchos a muchos
	dbHelper.DB.Model(&oldUser).Updates(&newInfo)
	dbHelper.DB.Model(&oldUser).Association("Rols").Replace(&newInfo.Rols)
	return oldUser
}

func (ur UserRepository) deleteUser(id uint) {
	dbHelper.DB.Delete(&UserModel{}, id)
}
