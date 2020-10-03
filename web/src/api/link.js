import service from '@/utils/request'

// @Tags Link
// @Summary 创建Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Link true "创建Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /link/createLink [post]
export const createLink = (data) => {
     return service({
         url: "/link/createLink",
         method: 'post',
         data
     })
 }


// @Tags Link
// @Summary 删除Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Link true "删除Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /link/deleteLink [delete]
 export const deleteLink = (data) => {
     return service({
         url: "/link/deleteLink",
         method: 'delete',
         data
     })
 }

// @Tags Link
// @Summary 删除Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /link/deleteLink [delete]
 export const deleteLinkByIds = (data) => {
     return service({
         url: "/link/deleteLinkByIds",
         method: 'delete',
         data
     })
 }

// @Tags Link
// @Summary 更新Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Link true "更新Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /link/updateLink [put]
 export const updateLink = (data) => {
     return service({
         url: "/link/updateLink",
         method: 'put',
         data
     })
 }


// @Tags Link
// @Summary 用id查询Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Link true "用id查询Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /link/findLink [get]
 export const findLink = (params) => {
     return service({
         url: "/link/findLink",
         method: 'get',
         params
     })
 }


// @Tags Link
// @Summary 分页获取Link列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Link列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /link/getLinkList [get]
 export const getLinkList = (params) => {
     return service({
         url: "/link/getLinkList",
         method: 'get',
         params
     })
 }