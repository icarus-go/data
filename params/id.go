package params

import (
	"github.com/icarus-go/data/resource"
	"gorm.io/gorm"
)

//IWhere
//  Author: Kevin·CC
//  Description: 条件
type IWhere interface {
	Scopes(name ...string) func(tx *gorm.DB) *gorm.DB
}

type (
	GetByID struct {
		ID uint64 `json:"id" validate:"required" form:"id" swaggertype:"string" example:"uint64 主键ID"` // 主键ID
	}

	//GetBySnowFlakeID 获取单条记录
	GetBySnowFlakeID struct {
		ID resource.SnowFlakeID `json:"id" validate:"required" swaggertype:"string" example:"string 雪花算法ID(实际为无符号整形)"` //  ID 主键
	}

	// GetBySnowFlakeIDs 获取ids
	GetBySnowFlakeIDs struct {
		IDs []resource.SnowFlakeID `json:"id" validate:"required" swaggertype:"string" example:"[]string 雪花算法ID(实际为无符号整形)"` //  ID 主键
	}

	GetByIDs struct {
		IDs []uint64 `json:"ids" form:"ids" validate:"required" swaggertype:"string" example:"[]uint64 主键IDs"` // 主键IDs
	}

	GetByUUID struct {
		ID string `json:"id" form:"id" validate:"required" example:"string UUID"` // ID UUID
	}

	GetByUUIDs struct {
		IDs []string `json:"ids" form:"ids" validate:"required" example:"[]string UUIDs"` // IDs UUIDs
	}
)

func (g GetByID) Scopes(name ...string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(idWhere(name...)+" = ?", g.ID)
	}
}

func (g GetBySnowFlakeID) Scopes(name ...string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(idWhere(name...)+" = ?", g.ID)
	}
}

func (g GetByUUID) Scopes(name ...string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(idWhere(name...)+" = ?", g.ID)
	}
}

func (g GetByIDs) Scopes(name ...string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(idWhere(name...)+" in ?", g.IDs)
	}
}

func (g GetBySnowFlakeIDs) Scopes(name ...string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(idWhere(name...)+" in ?", g.IDs)
	}
}

func (g GetByUUIDs) Scopes(name ...string) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where(idWhere(name...)+" in ?", g.IDs)
	}
}

// idWhere
//  Author: Kevin·CC
//  Description: id过滤
//  Param name
//  Return string
func idWhere(name ...string) string {
	if len(name) > 0 {
		return name[0]
	}
	return "id"
}
