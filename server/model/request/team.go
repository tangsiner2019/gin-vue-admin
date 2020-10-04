package request

import "gin-vue-admin/model"

type TeamSearch struct {
	model.Team
	PageInfo
}
