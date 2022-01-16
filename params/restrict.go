package params

import "gorm.io/gorm"

//IRestrict
//  Author: Kevin·CC
//  Description: 限制接口
type IRestrict interface {
	Where() func(tx *gorm.DB) *gorm.DB
}

//RestrictPreload
//  Author: Kevin·CC
//  Description: 限制预加载
type RestrictPreload struct {
	Object string `json:"object"`
	IRestrict
}

// NewRestrict
//  Author: Kevin·CC
//  Description: 构建限制Preload
//  Param relationMapping 关系映射
//  Param restrict 限制筛选
//  Return RestrictPreload
func NewRestrict(relationMapping string, restrict IRestrict) RestrictPreload {
	return RestrictPreload{
		Object:    relationMapping,
		IRestrict: restrict,
	}
}

// NewRelation
//  Author: Kevin·CC
//  Description: 无条件加载关系
//  Param relation 关系名称
//  Return RestrictPreload
func NewRelation(relation string) RestrictPreload {
	return RestrictPreload{
		Object: relation,
	}
}

// NewRelations
//  Author: Kevin·Cai
//  Description: 无条件加载关系
//  Param relations 关系列表
//  Return entities 关系预加载列表
func NewRelations(relations ...string) (entities []RestrictPreload) {
	for _, relation := range relations {
		entities = append(entities, NewRelation(relation))
	}
	return
}
