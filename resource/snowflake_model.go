package resource

import (
	"gorm.io/gorm"
	"time"
)

//SnowflakeModel
//  Author: Kevin·CC
//  Description: 支持雪花算法的基础model
type SnowflakeModel struct {
	ID        SnowFlakeID    `json:"id" gorm:"column:id;size:20;primaryKey" swaggertype:"string" example:"string (实际为雪花算法ID) 主键ID"` // ID
	CreatedAt Datetime       `json:"createdAt,omitempty" swaggertype:"string" example:"创建时间"`                                       // CreatedAt 创建时间
	UpdatedAt Datetime       `json:"updatedAt,omitempty" swaggertype:"string" example:"更新时间"`                                       // UpdatedAt 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`                                                           // DeletedAt 删除时间
}

//BeforeCreate
//  Author: Kevin·CC
//  Description: 创建前设置值
//  Param tx 事务
//  Return error 错误信息
func (m *SnowflakeModel) BeforeCreate(*gorm.DB) error {
	now := time.Now().Local()
	m.CreatedAt.Time = now
	m.UpdatedAt.Time = now
	return nil
}

//BeforeUpdate
//  Author: Kevin·CC
//  Description: 更新前设置值
//  Param *gorm.DB 查询语句
//  Return error 错误信息
func (m *SnowflakeModel) BeforeUpdate(*gorm.DB) error {
	now := time.Now().Local()
	m.UpdatedAt.Time = now
	return nil
}
