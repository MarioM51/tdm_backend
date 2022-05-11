package blog

import (
	"users_api/src/errorss"
)

type IBlogService interface {
	findAll() *[]BlogModel
	findById(id int) *BlogModel
	save(newBlog *BlogModel) *BlogModel
	update(newBlog *BlogModel) *BlogModel
	deleteById(idToDel int) *BlogModel

	addLike(idBlog int, IdUser int) int
	removeLike(idBlog int, IdUser int) int
}

type BlogService struct {
}

var blogRepo IBlogRepository = BlogRepository{}

const _ANON_USER_ID = 1

func (BlogService) findAll() *[]BlogModel {
	all := blogRepo.findAll()
	return all
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

	errMsg := newInfoBlog.validate()
	if errMsg != "" {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: errMsg})
	}

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
