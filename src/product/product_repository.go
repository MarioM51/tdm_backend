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
	dbHelper.DB.AutoMigrate(&LikeProduct{})
}

func (pr ProductRepository) findAll() *[]ProductModel {
	allProducts := []ProductModel{}
	dbHelper.DB.Find(&allProducts)

	//add image struct only with the id
	var ids []int
	for i := range allProducts {
		ids = append(ids, allProducts[i].ID)
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

	//add num of likes

	for i := range allProducts {
		likes := pr.findAllLikesOfProduct(allProducts[i].ID)
		countLikes := len(likes)
		allProducts[i].Likes = countLikes
	}

	return &allProducts
}

func (ProductRepository) save(toSave *ProductModel) *ProductModel {
	//dbHelper.DB.Create(toSave)
	toSave.Image = ProductImage{}
	dbHelper.DB.Omit(clause.Associations).Create(toSave)
	return toSave
}

func (pr ProductRepository) findById(id int) (finded *ProductModel) {
	finded = &ProductModel{}
	dbHelper.DB.Find(finded, id)
	if finded.ID <= 0 {
		return nil
	}

	likes := pr.findAllLikesOfProduct(finded.ID)
	countLikes := len(likes)
	finded.Likes = countLikes

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

//=========images

func (ProductRepository) findImageByProductId(id int) *ProductImage {
	productFinded := &ProductModel{}
	dbHelper.DB.Preload(_IMAGE_ASO).Find(productFinded, id)
	if productFinded.ID <= 0 || productFinded.Image.ID <= 0 {
		return nil
	}
	return &productFinded.Image
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

//=========likes

func (ProductRepository) findUserLikes(idUser int) []LikeProduct {
	finds := []LikeProduct{}
	dbHelper.DB.Where("fk_user = ?", idUser).Find(&finds)
	return finds
}

func (ProductRepository) findAllLikesOfProduct(idProduct int) (allLikes []LikeProduct) {
	dbHelper.DB.Where("fk_product = ?", idProduct).Find(&allLikes)
	return allLikes
}

func (ps ProductRepository) addLike(idProduct int, idUser int) int {
	toSave := &LikeProduct{
		FkProduct: idProduct,
		FKUser:    idUser,
	}
	dbHelper.DB.Create(toSave)

	allLikes := ps.findAllLikesOfProduct(toSave.FkProduct)
	return len(allLikes)
}

func (ps ProductRepository) removeLike(idProduct int, idUser int) int {
	toDel := LikeProduct{}
	dbHelper.DB.Where("fk_product = ? AND fk_user = ?", idProduct, idUser).First(&toDel)

	if toDel.FkProduct >= 1 {
		dbHelper.DB.Where("created_at = ?", toDel.CreatedAt).Delete(&toDel)
	}

	allLikes := ps.findAllLikesOfProduct(idProduct)
	return len(allLikes)
}
