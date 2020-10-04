package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateProduct
// @description   create a Product
// @param     product               model.Product
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateProduct(product model.Product) (err error) {
	err = global.GVA_DB.Create(&product).Error
	return err
}

// @title    DeleteProduct
// @description   delete a Product
// @auth                     （2020/04/05  20:22）
// @param     product               model.Product
// @return                    error

func DeleteProduct(product model.Product) (err error) {
	err = global.GVA_DB.Delete(product).Error
	return err
}

// @title    DeleteProductByIds
// @description   delete Products
// @auth                     （2020/04/05  20:22）
// @param     product               model.Product
// @return                    error

func DeleteProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Product{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateProduct
// @description   update a Product
// @param     product          *model.Product
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateProduct(product *model.Product) (err error) {
	err = global.GVA_DB.Save(product).Error
	return err
}

// @title    GetProduct
// @description   get the info of a Product
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Product        Product

func GetProduct(id uint) (err error, product model.Product) {
	err = global.GVA_DB.Where("id = ?", id).First(&product).Error
	return
}

// @title    GetProductInfoList
// @description   get Product list by pagination, 分页获取Product
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetProductInfoList(info request.ProductSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Product{})
	var products []model.Product
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&products).Error
	return err, products, total
}
