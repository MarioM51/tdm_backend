package product

import "users_api/src/helpers"

type ProductRepository struct{}

var dbHelper = helpers.DBHelper{}

func CreateProductSchema() {
	dbHelper.Connect()

	dbHelper.DB.AutoMigrate(&ProductModel{})
}

func (ProductRepository) findAll() (all *[]ProductModel) {
	all = &[]ProductModel{}
	dbHelper.DB.Find(all)
	return all
}

func (ProductRepository) save(toSave *ProductModel) *ProductModel {
	dbHelper.DB.Create(toSave)
	return toSave
}

func (ProductRepository) findById(id int) (finded *ProductModel) {
	finded = &ProductModel{}
	dbHelper.DB.Find(finded, id)
	return finded
}

func (ProductRepository) delete(toDelete *ProductModel) *ProductModel {
	dbHelper.DB.Delete(toDelete)
	return toDelete
}

func (ProductRepository) update(oldInfo, newInfo *ProductModel) *ProductModel {
	dbHelper.DB.Model(&oldInfo).Updates(&newInfo)
	return newInfo
}
