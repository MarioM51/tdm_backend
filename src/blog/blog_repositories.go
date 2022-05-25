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

	findUserBlogLikes(idUser int) []LikeBlog
	addLike(idBlog int, iduser int) int
	removeLike(idProduct int, idUser int) int
}

type BlogRepository struct {
}

var dbHelper = helpers.DBHelper{}

func CreateBlogSchema() {
	dbHelper.Connect()

	/*
		dbHelper.DB.AutoMigrate(&BlogModel{})
		dbHelper.DB.AutoMigrate(&LikeBlog{})
	*/
}

func (br BlogRepository) findAll() *[]BlogModel {
	all := []BlogModel{}
	dbHelper.DB.Select("id", "title", "author", "abstract", "created_at", "updated_at").Find(&all)
	for i := range all {
		likes := br.findAllLikesOfBlog(all[i].Id)
		likesCount := len(likes)
		all[i].Likes = likesCount
	}
	return &all
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

func (br BlogRepository) findById(id int) *BlogModel {
	finded := BlogModel{}
	dbHelper.DB.Find(&finded, id)

	if finded.Id == 0 {
		return nil
	}

	likes := br.findAllLikesOfBlog(finded.Id)
	likesCount := len(likes)
	finded.Likes = likesCount

	return &finded
}

func handleTxError(txErr error) {
	isUserFault := strings.Contains(strings.ToLower(txErr.Error()), `unique constraint`)
	if isUserFault {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "That title already exist, try a different title"})
	}
	panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error saving blog"})
}

// likes

func (BlogRepository) findUserBlogLikes(idUser int) []LikeBlog {
	finds := []LikeBlog{}
	dbHelper.DB.Where("fk_user = ?", idUser).Find(&finds)
	return finds
}

func (BlogRepository) findAllLikesOfBlog(idBlog int) (allLikes []LikeBlog) {
	dbHelper.DB.Where("fk_blog = ?", idBlog).Find(&allLikes)
	return allLikes
}

func (ps BlogRepository) addLike(idBlog int, idUser int) int {
	toSave := &LikeBlog{
		FkBlog: idBlog,
		FKUser: idUser,
	}
	dbHelper.DB.Create(toSave)

	allLikes := ps.findAllLikesOfBlog(toSave.FkBlog)
	return len(allLikes)
}

func (ps BlogRepository) removeLike(idBlog int, idUser int) int {
	toDel := LikeBlog{}
	dbHelper.DB.Where("fk_blog = ? AND fk_user = ?", idBlog, idUser).First(&toDel)

	if toDel.FkBlog >= 1 {
		dbHelper.DB.Where("created_at = ?", toDel.CreatedAt).Delete(&toDel)
	}

	allLikes := ps.findAllLikesOfBlog(idBlog)
	return len(allLikes)
}
