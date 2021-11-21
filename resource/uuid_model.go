package resource

import (
	"gorm.io/gorm"
	"time"
)

//UUIDModel
//  Author: Kevin·CC
//  Description:
type UUIDModel struct {
	ID        string         `json:"id" gorm:"column:id;type:varchar(32);primaryKey" swaggertype:"string" example:"string UUIDModel 主键ID"` // ID
	CreatedAt Datetime       `json:"createdAt,omitempty" swaggertype:"string" example:"创建时间"`                                              // CreatedAt 创建时间
	UpdatedAt Datetime       `json:"updatedAt,omitempty" swaggertype:"string" example:"更新时间"`                                              // UpdatedAt 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`                                                                  // DeletedAt 删除时间
}

//BeforeCreate
//  Author: Kevin·CC
//  Description: 创建钩子
//  Param tx
//  Return error 错误信息
func (s *UUIDModel) BeforeCreate(tx *gorm.DB) error {
	now := time.Now().Local()
	s.CreatedAt.Time = now
	s.UpdatedAt.Time = now
	return nil
}

//BeforeUpdate
//  Author: Kevin·CC
//  Description: 更新前钩子
//  Param tx 事务
//  Return error 错误信息
func (s *UUIDModel) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now().Local()
	s.UpdatedAt.Time = now
	return nil
}
