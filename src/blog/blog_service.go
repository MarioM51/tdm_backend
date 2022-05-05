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
}

type BlogService struct {
}

var blogRepo IBlogRepository = BlogRepository{}

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
