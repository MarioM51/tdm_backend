package product

import (
	"users_api/src/helpers"

	"gorm.io/gorm/clause"
)

type ProductRepository struct{}

var dbHelper = helpers.DBHelper{}

const _IMAGE_ASO = "Image"

func CreateProductSchema() {
	dbHelper.Connect()

	dbHelper.DB.AutoMigrate(&ProductImage{})
	dbHelper.DB.AutoMigrate(&ProductModel{})
}

func (ProductRepository) findAll() *[]ProductModel {
	allProducts := []ProductModel{}
	dbHelper.DB.Find(&allProducts)

	var ids []int
	for _, v := range allProducts {
		ids = append(ids, v.ID)
	}
	allImages := []ProductImage{}
	dbHelper.DB.Select("id, updated_at").Find(&allImages, ids)

	for iProduct := 0; iProduct < len(allProducts); iProduct++ {
		for iImage := 0; iImage < len(allImages); iImage++ {
			if allProducts[iProduct].ID == allImages[iImage].ID {
				allProducts[iProduct].Image = allImages[iImage]
			}
		}
	}

	return &allProducts
}

func (ProductRepository) save(toSave *ProductModel) *ProductModel {
	//dbHelper.DB.Create(toSave)
	toSave.Image = ProductImage{}
	dbHelper.DB.Omit(clause.Associations).Create(toSave)
	return toSave
}

func (ProductRepository) saveImage(idProduct int, newImage *ProductImage) *ProductImage {
	image := ProductImage{}
	newImage.ID = idProduct
	dbHelper.DB.Find(&image, idProduct)
	if image.ID > 0 {
		dbHelper.DB.Delete(&image, idProduct)
	}
	dbHelper.DB.Create(newImage)
	newImage.Base64 = ""
	return newImage
}

func (ProductRepository) findById(id int) (finded *ProductModel) {
	finded = &ProductModel{}
	dbHelper.DB.Find(finded, id)
	return finded
}

func (ProductRepository) findImageByProductId(id int) *ProductImage {
	productFinded := &ProductModel{}
	dbHelper.DB.Preload(_IMAGE_ASO).Find(productFinded, id)
	if productFinded.ID <= 0 || productFinded.Image.ID <= 0 {
		return nil
	}
	return &productFinded.Image
}

func (ProductRepository) delete(toDelete *ProductModel) *ProductModel {
	dbHelper.DB.Delete(toDelete)
	return toDelete
}

func (ProductRepository) update(oldInfo, newInfo *ProductModel) *ProductModel {
	dbHelper.DB.Model(&oldInfo).Updates(&newInfo)
	return newInfo
}
