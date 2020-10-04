package request

import "gin-vue-admin/model"

type NewsSearch struct {
	model.News
	PageInfo
}
