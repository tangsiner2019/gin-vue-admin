package request

import "gin-vue-admin/model"

type BannerSearch struct {
	model.Banner
	PageInfo
}
