package blog

import (
	"strings"
	"users_api/src/errorss"
)

type IBlogRepository interface {
	findAll(filter string) *[]BlogModel
	save(newBlog *BlogModel) *BlogModel
	update(oldInfo, newInfo *BlogModel) *BlogModel
	findById(id int) *BlogModel
	delete(toDel *BlogModel) *BlogModel

	findUserBlogLikes(idUser int) []LikeBlog
	addLike(idBlog int, iduser int) int
	removeLike(idProduct int, idUser int) int

	findBlogComment(idComment int) *BlogComment
	addComment(toAdd BlogComment) BlogComment
	deleteComment(idComment int)
}

type BlogRepository struct {
}

func (br BlogRepository) findAll(filter string) *[]BlogModel {
	all := []BlogModel{}
	dbHelper.DB.
		Select("id", "title", "author", "abstract", "created_at", "updated_at").
		Where(filter).
		Find(&all)
	for i := range all {
		likes := br.findAllLikesOfBlog(all[i].Id)
		likesCount := len(likes)
		all[i].Likes = likesCount
	}

	for i := range all {
		comments := br.findAllCommentsOfBlog(all[i].Id)

		commentsAmount := len(comments)
		all[i].CommentCount = commentsAmount

		//get average rating
		var total int = 0
		for j := range comments {
			total = total + comments[j].Rating
		}

		if total == 0 && commentsAmount == 0 {
			all[i].CommentsRating = 0
		} else {
			all[i].CommentsRating = float32(total) / float32(commentsAmount)
		}

	}

	return &all
}

func (BlogRepository) save(newBlog *BlogModel) *BlogModel {
	omits := []string{}
	if newBlog.OnHomeScreen.Year() <= 1 {
		omits = append(omits, "on_home_screen")
	}

	tx := dbHelper.DB.Omit(omits...).Create(newBlog)
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

	if newInfo.OnHomeScreen.Year() == 1 {
		dbHelper.DB.Model(&oldInfo).Update("on_home_screen", nil)
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

	comments := br.findAllCommentsOfBlog(finded.Id)
	finded.Comments = comments

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

//comments
func (BlogRepository) findAllCommentsOfBlog(idBlog int) (allComments []BlogComment) {
	dbHelper.DB.Where("id_blog = ?", idBlog).Find(&allComments)
	return allComments
}

func (BlogRepository) addComment(toAdd BlogComment) BlogComment {
	tx := dbHelper.DB.Create(&toAdd)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error saving comment blog"})
	}

	return toAdd
}

func (BlogRepository) findBlogComment(idComment int) (finded *BlogComment) {
	dbHelper.DB.First(&finded, idComment)
	if finded.Id == 0 {
		return nil
	}

	return finded
}

func (BlogRepository) deleteComment(idComment int) {
	tx := dbHelper.DB.Delete(&BlogComment{}, idComment)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error deleting blog comment"})
	}
}
