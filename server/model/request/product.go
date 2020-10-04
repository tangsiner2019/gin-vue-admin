package request

import "gin-vue-admin/model"

type ProductSearch struct {
	model.Product
	PageInfo
}
