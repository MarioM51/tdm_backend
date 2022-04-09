package blog

import (
	"strings"
	"users_api/src/errorss"
	"users_api/src/helpers"
)

type IBlogRepository interface {
	findAll() *[]BlogModel
	save(newBlog *BlogModel) *BlogModel
	update(oldInfo, newInfo *BlogModel) *BlogModel
	findById(id int) *BlogModel
	delete(toDel *BlogModel) *BlogModel
}

type BlogRepository struct {
}

var dbHelper = helpers.DBHelper{}

func CreateBlogSchema() {
	dbHelper.Connect()

	dbHelper.DB.AutoMigrate(&BlogModel{})
}

func (BlogRepository) findAll() *[]BlogModel {
	all := &[]BlogModel{}
	//dbHelper.DB.Select("id", "title", "autor", "created_at", "updated_at").Find(&all)
	dbHelper.DB.Find(&all)
	return all
}

func (BlogRepository) save(newBlog *BlogModel) *BlogModel {
	tx := dbHelper.DB.Create(newBlog)
	if tx.Error != nil {
		handleTxError(tx.Error)
	}

	return newBlog
}

func (BlogRepository) update(oldInfo, newInfo *BlogModel) *BlogModel {
	tx := dbHelper.DB.Model(&oldInfo).Updates(&newInfo)

	if tx.Error != nil {
		handleTxError(tx.Error)
	}

	return oldInfo
}

func (BlogRepository) delete(toDel *BlogModel) *BlogModel {
	dbHelper.DB.Delete(toDel, toDel.Id)
	return toDel
}

func (BlogRepository) findById(id int) *BlogModel {
	finded := &BlogModel{}
	dbHelper.DB.Find(&finded, id)

	if finded.Id == 0 {
		return nil
	}

	return finded
}

func handleTxError(txErr error) {
	isUserFault := strings.Contains(strings.ToLower(txErr.Error()), `unique constraint`)
	if isUserFault {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "That title already exist, try a different title"})
	}
	panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error saving blog"})
}
