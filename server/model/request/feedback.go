package request

import "gin-vue-admin/model"

type FeedbackSearch struct {
	model.Feedback
	PageInfo
}
