package request

import "gin-vue-admin/model"

type ServerSearch struct {
	model.Server
	PageInfo
}
