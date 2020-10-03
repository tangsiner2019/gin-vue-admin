package request

import "gin-vue-admin/model"

type LinkSearch struct{
    model.Link
    PageInfo
}