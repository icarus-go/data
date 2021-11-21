package resource

import (
	"gorm.io/gorm"
	"time"
)

//Model
//  Author: Kevin·CC
//  Description: 自增IDmodel
type Model struct {
	ID        uint64         `json:"id" gorm:"column:id;type:bigin(20);primaryKey" swaggertype:"string" example:"uint64 主键ID"` // ID
	CreatedAt Datetime       `json:"createdAt,omitempty" swaggertype:"string" example:"创建时间"`                                  // 创建时间
	UpdatedAt Datetime       `json:"updatedAt,omitempty" swaggertype:"string" example:"更新时间"`                                  // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`                                                      // 删除时间
}

//BeforeCreate
//  Author: Kevin·CC
//  Description: 创建钩子
//  Param tx 事务
//  Return error 错误信息
func (m *Model) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().Local()
	m.CreatedAt.Time = now
	m.UpdatedAt.Time = now
	return nil
}

//BeforeUpdate
//  Author: Kevin·CC
//  Description: 更新钩子
//  Param tx 事务
//  Return error 错误信息
func (m *Model) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now().Local()
	m.UpdatedAt.Time = now
	return nil
}
