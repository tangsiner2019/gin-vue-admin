package request

import "gin-vue-admin/model"

type CateSearch struct {
	model.Cate
	PageInfo
}
