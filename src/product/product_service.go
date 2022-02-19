package product

import "users_api/src/errorss"

type ProductService struct{}

var productDao = ProductRepository{}

func (ProductService) findAll() (all *[]ProductModel) {
	all = productDao.findAll()
	return all
}

func (ProductService) save(toSave *ProductModel) (saved *ProductModel) {
	saved = productDao.save(toSave)
	return saved
}

func (ProductService) findById(id int) (finded *ProductModel) {
	finded = productDao.findById(id)
	if finded == nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 404, Cause: "Product not found"})
	}

	return finded
}

func (ps ProductService) update(newInfo *ProductModel) (updated *ProductModel) {
	if newInfo.Id <= 0 {
		panic(errorss.ErrorResponseModel{HttpStatus: 400, Cause: "id required and must be greater than 0"})
	}

	oldInfo := ps.findById(newInfo.Id)

	updated = productDao.update(oldInfo, newInfo)
	return updated
}

func (ps ProductService) delete(id int) (deleted *ProductModel) {
	toDel := ps.findById(id)

	deleted = productDao.delete(toDel)
	return deleted
}
