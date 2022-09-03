package product

import (
	"users_api/src/errorss"
)

type ProductService struct{}

var productDao = ProductRepository{}

const _ANON_USER_ID = 1

func (ProductService) findAll() (all *[]ProductModel) {
	all = productDao.findAll("")
	return all
}

func (ProductService) FindOnHomeScreen() (all *[]ProductModel) {
	all = productDao.findAll("on_home_screen IS NOT NULL")
	return all
}

func (ProductService) save(toSave *ProductModel) (saved *ProductModel) {
	saved = productDao.save(toSave)
	if saved.ID == 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error saving product"})
	}
	return saved
}

func (ps ProductService) saveImage(idProduct int, toSave *ProductImage) (saved *ProductImage) {
	productFinded := ps.findById(idProduct)
	saved = productDao.saveImage(productFinded.ID, toSave)
	return saved
}

func (ProductService) findById(id int) (finded *ProductModel) {
	finded = productDao.findById(id)
	if finded == nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 404, Cause: "Product not found"})
	}

	return finded
}

func (ProductService) findImageIdImage(idImage int) (finded *ProductImage) {
	finded = productDao.findImageByIdImage(idImage)
	if finded == nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 404, Cause: "Product Image not found"})
	}

	return finded
}

func (ProductService) deleteImageIdImage(idImage int) (finded *ProductImage) {
	finded = productDao.deleteImageIdImage(idImage)
	if finded == nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 404, Cause: "Product Image not found"})
	}

	return finded
}

func (ps ProductService) update(newInfo *ProductModel) (updated *ProductModel) {
	if newInfo.ID <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "id required and must be greater than 0"})
	}

	oldInfo := ps.findById(newInfo.ID)

	// avoid updates not wanted
	newInfo.ID = 0
	newInfo.Images = nil
	newInfo.Likes = 0
	updated = productDao.update(oldInfo, newInfo)

	return updated
}

func (ps ProductService) delete(id int) (deleted *ProductModel) {
	toDel := ps.findById(id)

	deleted = productDao.delete(toDel)
	return deleted
}

//==== likes

func (ps ProductService) addLike(idProduct int, idUser int) int {
	if idUser <= 0 {
		// we change to user 1 that is the anonymous user
		idUser = _ANON_USER_ID
	}

	finded := ps.findById(idProduct) // panic if not exists

	if idUser != _ANON_USER_ID {
		finds := productDao.findUserLikes(idUser)
		for _, like := range finds {
			if like.FkProduct == idProduct {
				panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "User already add like to this product"})
			}
		}
	}

	likesCount := productDao.addLike(finded.ID, idUser)
	return likesCount
}

func (ps ProductService) removeLike(idProduct int, idUser int) int {
	if idUser <= 0 {
		idUser = _ANON_USER_ID
	}
	likesCount := productDao.removeLike(idProduct, idUser)
	return likesCount
}

func (bs ProductService) addComment(toAdd Comment) Comment {
	bs.findById(toAdd.IdTarget)

	commentAdded := productDao.addComment(toAdd)

	return commentAdded
}

func (bs ProductService) deleteComment(toDel Comment) Comment {
	bs.findById(toDel.IdTarget)

	finded := productDao.findProductComment(toDel.Id)
	if finded == nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error deleting blog comment"})
	}

	if toDel.IdUser != 666777 {
		if finded.IdUser != toDel.IdUser {
			panic(errorss.ErrorResponseModel{HttpStatus: 403, Cause: "The comment doesn`t belong to the user"})
		}
	}

	productDao.deleteComment(toDel.Id)

	return *finded
}
