package users

import (
	"users_api/src/helpers"

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

var dbHelper = helpers.DBHelper{}

func CreateUserSchema() {
	dbHelper.Connect()

	/*
		dbHelper.DB.AutoMigrate(&RoleModel{})
		dbHelper.DB.AutoMigrate(&UserModel{})

		adminRole := RoleModel{Model: gorm.Model{ID: 79}, Name: "admin"}
		adminUser := UserModel{Model: gorm.Model{ID: 79}, Email: "mario2@email.com", Rols: []RoleModel{{Model: gorm.Model{ID: 79}}}, Password: "$2a$12$OenFL4B1HRFZasAuL2my5.PNJ2GRR4wLl1BUDH2vl0ZBeU2Dv3.Gq"}
		dbHelper.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&adminRole)
		dbHelper.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&adminUser)

		anonUser := UserModel{Model: gorm.Model{ID: 1}, Email: "anon@email.com", Password: "$2a$12$OenFL4B1HRFZasAuL2my5.PNJ2GRR4wLl1BUDH2vl0ZBeU2Dv3.Gq"}
		dbHelper.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&anonUser)
	*/
}

func (ur UserRepository) saveUser(newUser *UserModel) *UserModel {
	newUser.Rols = nil
	dbHelper.DB.Create(&newUser)
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
