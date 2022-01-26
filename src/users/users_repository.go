package users

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&RoleModel{})
	db.AutoMigrate(&UserModel{})

	adminRole := RoleModel{Model: gorm.Model{ID: 79}, Name: "admin"}
	adminUser := UserModel{Model: gorm.Model{ID: 79}, Email: "mario2@email.com", Roles: []RoleModel{{Model: gorm.Model{ID: 79}}}, Password: "$2a$12$OenFL4B1HRFZasAuL2my5.PNJ2GRR4wLl1BUDH2vl0ZBeU2Dv3.Gq"}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&adminRole)
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&adminUser)
}

func (ur UserRepository) saveUser(newUser *UserModel) *UserModel {
	newUser.Roles = nil
	db.Create(&newUser)
	return newUser
}

func (ur UserRepository) findAll() (allUsers *[]UserModel) {
	db.Find(&allUsers)
	for i := range *allUsers {
		(*allUsers)[i].Password = ""
	}
	return allUsers
}

func (_ UserRepository) findUserById(id uint) *UserModel {
	var userFinded *UserModel
	db.Preload(clause.Associations).Find(&userFinded, id)
	if userFinded.ID <= 0 {
		return nil
	}
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
	newInfo.Roles = nil
	db.Model(&oldUser).Updates(newInfo)
	return oldUser
}

func (ur UserRepository) deleteUser(id uint) {
	db.Delete(&UserModel{}, id)
}
