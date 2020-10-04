package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateFeedback
// @description   create a Feedback
// @param     feedback               model.Feedback
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateFeedback(feedback model.Feedback) (err error) {
	err = global.GVA_DB.Create(&feedback).Error
	return err
}

// @title    DeleteFeedback
// @description   delete a Feedback
// @auth                     （2020/04/05  20:22）
// @param     feedback               model.Feedback
// @return                    error

func DeleteFeedback(feedback model.Feedback) (err error) {
	err = global.GVA_DB.Delete(feedback).Error
	return err
}

// @title    DeleteFeedbackByIds
// @description   delete Feedbacks
// @auth                     （2020/04/05  20:22）
// @param     feedback               model.Feedback
// @return                    error

func DeleteFeedbackByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Feedback{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateFeedback
// @description   update a Feedback
// @param     feedback          *model.Feedback
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateFeedback(feedback *model.Feedback) (err error) {
	err = global.GVA_DB.Save(feedback).Error
	return err
}

// @title    GetFeedback
// @description   get the info of a Feedback
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Feedback        Feedback

func GetFeedback(id uint) (err error, feedback model.Feedback) {
	err = global.GVA_DB.Where("id = ?", id).First(&feedback).Error
	return
}

// @title    GetFeedbackInfoList
// @description   get Feedback list by pagination, 分页获取Feedback
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetFeedbackInfoList(info request.FeedbackSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Feedback{})
	var feedbacks []model.Feedback
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&feedbacks).Error
	return err, feedbacks, total
}
