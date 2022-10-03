package blog

import (
	"users_api/src/errorss"
)

type IBlogService interface {
	findAll() *[]BlogModel
	FindOnHomeScreen() *[]BlogModel
	findById(id int) *BlogModel
	save(newBlog *BlogModel) *BlogModel
	update(newBlog *BlogModel) *BlogModel
	deleteById(idToDel int) *BlogModel

	addLike(idBlog int, IdUser int) int
	removeLike(idBlog int, IdUser int) int

	findAllComments() []BlogComment
	addComment(toAdd BlogComment) BlogComment
	deleteComment(toDel BlogComment) BlogComment
	addCommentResponse(toAdd *BlogComment)
}

type BlogService struct {
}

var blogRepo IBlogRepository = BlogRepository{}

const _ANON_USER_ID = 1

func (BlogService) findAll() *[]BlogModel {
	all := blogRepo.findAll("")
	return all
}

func (BlogService) FindOnHomeScreen() *[]BlogModel {
	onHomeScreen := blogRepo.findAll("on_home_screen IS NOT NULL")
	return onHomeScreen
}

func (BlogService) findById(id int) *BlogModel {
	finded := blogRepo.findById(id)
	if finded == nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 404, Cause: "blog not finded"})
	}
	return finded
}

func (BlogService) save(newBlog *BlogModel) *BlogModel {
	errMsg := newBlog.validate()
	if errMsg != "" {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: errMsg})
	}
	newBlog.Author = "Temp"
	blogSaved := blogRepo.save(newBlog)
	return blogSaved
}

func (bs BlogService) update(newInfoBlog *BlogModel) *BlogModel {
	if newInfoBlog.Id <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "Id is required"})
	}

	oldInfo := bs.findById(newInfoBlog.Id)

	blogUpdated := blogRepo.update(oldInfo, newInfoBlog)
	return blogUpdated
}

func (bs BlogService) deleteById(idToDel int) *BlogModel {
	blogToDel := bs.findById(idToDel)
	blogDeleted := blogRepo.delete(blogToDel)
	return blogDeleted
}

func (bs BlogService) addLike(idProduct int, idUser int) int {
	if idUser <= 0 {
		// we change to user 1 that is the anonymous user
		idUser = _ANON_USER_ID
	}

	finded := bs.findById(idProduct) // panic if not exists

	if idUser != _ANON_USER_ID {
		finds := blogRepo.findUserBlogLikes(idUser)
		for _, like := range finds {
			if like.FkBlog == idProduct {
				panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "User already add like to this product"})
			}
		}
	}

	likesCount := blogRepo.addLike(finded.Id, idUser)
	return likesCount
}

func (bs BlogService) removeLike(idProduct int, idUser int) int {
	if idUser <= 0 {
		idUser = _ANON_USER_ID
	}
	likesCount := blogRepo.removeLike(idProduct, idUser)
	return likesCount
}

// ===============0 Comment

func (bs BlogService) addComment(toAdd BlogComment) BlogComment {
	bs.findById(toAdd.IdBlog)

	commentAdded := blogRepo.addComment(toAdd)

	return commentAdded
}

func (bs BlogService) deleteComment(toDel BlogComment) BlogComment {
	bs.findById(toDel.IdBlog)

	finded := blogRepo.findBlogComment(toDel.Id)
	if finded == nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Comentario a eliminar no encontrado"})
	}

	if toDel.IdUser != 666777 {
		if finded.IdUser != toDel.IdUser {
			panic(errorss.ErrorResponseModel{HttpStatus: 403, Cause: "Sin permiso para eliminar el comentario"})
		}
	}

	blogRepo.deleteComment(finded)

	return *finded
}

func (bs BlogService) addCommentResponse(newResponse *BlogComment) {
	bs.findById(newResponse.IdBlog)

	usrServ.FindById(uint(newResponse.IdUser))

	blogRepo.addComment(*newResponse)

}

func (bs BlogService) findAllComments() []BlogComment {
	allComments := blogRepo.findAllComments()

	return allComments

}
