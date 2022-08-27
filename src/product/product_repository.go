package product

import (
	"users_api/src/errorss"

	"gorm.io/gorm/clause"
)

type ProductRepository struct{}

const _IMAGE_ASO = "Image"

func (pr ProductRepository) findAll() *[]ProductModel {
	allProducts := []ProductModel{}
	dbHelper.DB.Find(&allProducts)

	//add image struct only with the id
	for i := range allProducts {
		pr.findProductImages(&allProducts[i])
	}

	//add num of likes

	for i := range allProducts {
		likes := pr.findAllLikesOfProduct(allProducts[i].ID)
		countLikes := len(likes)
		allProducts[i].Likes = countLikes
	}

	return &allProducts
}

func (ProductRepository) findProductImages(p *ProductModel) {
	//dbHelper.DB.Create(toSave)
	allImages := []ProductImage{}
	dbHelper.DB.Select("id, updated_at").Where("fk_product = ?", p.ID).Find(&allImages)
	p.Images = allImages
}

func (ProductRepository) save(toSave *ProductModel) *ProductModel {
	//dbHelper.DB.Create(toSave)
	toSave.Images = []ProductImage{}
	dbHelper.DB.Omit(clause.Associations).Create(toSave)
	return toSave
}

func (pr ProductRepository) findById(id int) (finded *ProductModel) {
	finded = &ProductModel{}
	dbHelper.DB.Find(finded, id)
	if finded.ID <= 0 {
		return nil
	}

	pr.findProductImages(finded)
	likes := pr.findAllLikesOfProduct(finded.ID)
	countLikes := len(likes)
	finded.Likes = countLikes

	finded.Comments = pr.findAllCommentsOfProduct(finded.ID)

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

func (ProductRepository) findImageByIdImage(idImage int) *ProductImage {
	image := &ProductImage{}
	dbHelper.DB.Preload(_IMAGE_ASO).Find(image, idImage)
	if image.ID <= 0 {
		return nil
	}
	return image
}

func (ProductRepository) saveImage(idProduct int, newImage *ProductImage) *ProductImage {
	/*
		image := ProductImage{}
		dbHelper.DB.Find(&image, idProduct)
		if image.ID > 0 {
			dbHelper.DB.Delete(&image, idProduct)
		}
	*/
	newImage.FkProduct = idProduct
	newImage.ID = 0
	tx := dbHelper.DB.Create(newImage)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error saving image product"})
	}
	newImage.Base64 = ""
	return newImage
}

func (ProductRepository) deleteImageIdImage(idImage int) (toDel *ProductImage) {
	finded := ProductImage{}
	dbHelper.DB.Select("id", "updated_at").First(&finded, idImage)
	if finded.ID <= 0 {
		return nil
	}

	tx := dbHelper.DB.Delete(&toDel, finded.ID)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error deleting image product"})
	}

	return &finded
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

//=========comment
func (ProductRepository) findAllCommentsOfProduct(idProduct int) (comments []Comment) {
	dbHelper.DB.Where("id_target = ?", idProduct).Find(&comments)
	return comments
}

func (ProductRepository) addComment(toAdd Comment) Comment {
	tx := dbHelper.DB.Create(&toAdd)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error saving comment product"})
	}

	return toAdd
}

func (ProductRepository) findProductComment(idComment int) (finded *Comment) {
	dbHelper.DB.First(&finded, idComment)
	if finded.Id == 0 {
		return nil
	}

	return finded
}

func (ProductRepository) deleteComment(idComment int) {
	tx := dbHelper.DB.Delete(&Comment{}, idComment)
	if tx.Error != nil {
		panic(errorss.ErrorResponseModel{HttpStatus: 500, Cause: "Error deleting product comment"})
	}
}
