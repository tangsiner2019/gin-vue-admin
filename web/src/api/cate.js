import service from '@/utils/request'

// @Tags Cate
// @Summary 创建Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cate true "创建Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cate/createCate [post]
export const createCate = (data) => {
     return service({
         url: "/cate/createCate",
         method: 'post',
         data
     })
 }


// @Tags Cate
// @Summary 删除Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cate true "删除Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cate/deleteCate [delete]
 export const deleteCate = (data) => {
     return service({
         url: "/cate/deleteCate",
         method: 'delete',
         data
     })
 }

// @Tags Cate
// @Summary 删除Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cate/deleteCate [delete]
 export const deleteCateByIds = (data) => {
     return service({
         url: "/cate/deleteCateByIds",
         method: 'delete',
         data
     })
 }

// @Tags Cate
// @Summary 更新Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cate true "更新Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cate/updateCate [put]
 export const updateCate = (data) => {
     return service({
         url: "/cate/updateCate",
         method: 'put',
         data
     })
 }


// @Tags Cate
// @Summary 用id查询Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cate true "用id查询Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cate/findCate [get]
 export const findCate = (params) => {
     return service({
         url: "/cate/findCate",
         method: 'get',
         params
     })
 }


// @Tags Cate
// @Summary 分页获取Cate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Cate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cate/getCateList [get]
 export const getCateList = (params) => {
     return service({
         url: "/cate/getCateList",
         method: 'get',
         params
     })
 }