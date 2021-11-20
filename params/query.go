package params

import (
	"strings"
)

//Paging
//  Author: Kevin·CC
//  Description: 分页
type Paging struct {
	Page     int `json:"page" form:"page" example:"1"`          // 页码
	PageSize int `json:"pageSize" form:"pageSize" example:"20"` // 页面最大条数
}

//ResetPage 将不合理的分页值进行重置
func (q *Paging) ResetPage() *Paging {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.PageSize < 1 {
		q.PageSize = 20
	}
	return q
}

//Order
//  Author: Kevin·CC
//  Description: 排序
type Order struct {
	Fields []string `json:"fields" form:"fields" example:"id,createTime"` // 排序字段
	Type   string   `json:"type" form:"type" example:"desc/asc"`          // 排序类型
}

func (o *Order) Join() string {
	return strings.Join(o.Fields, ",")
}

func (o *Order) Sort() string {
	if o.Type == "" {
		return ""
	} // 如果为空，默认使用降序
	if o.Type == "desc" || o.Type == "asc" {
		return o.Type
	} // 符合数据库值，返回
	return "" // 不符合，直接返回空，按照升序处理

}
