package resource

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
)

// Boolean 类似Java概念中的boolean 从tinyint(1)转为boolean, 代码更为直观 省去if/else转换值
type Boolean bool

// Int 布尔对应数值
//  Author:  Kevin·CC
func (b Boolean) Int() int {
	if b {
		return 1
	}
	return 0
}

// String
//  Description: 字符串值
//  Author: Kevin·CC
//  return string
func (b Boolean) String() string {
	if b {
		return "true"
	}
	return "false"
}

// Scan 扫描
// Author Kevin·CC
func (b *Boolean) Scan(value interface{}) error {
	nullBool := sql.NullBool{}
	err := nullBool.Scan(value)
	*b = Boolean(nullBool.Bool)
	return err
}

// Value 值
// Author Kevin·CC
func (b *Boolean) Value() (driver.Value, error) {
	return driver.Value(bool(*b)), nil
}

// MarshalJSON 序列化
// Author Kevin·CC
func (b *Boolean) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`%v`, *b)), nil
}

// UnmarshalJSON 反序列化
// Author Kevin·CC
func (b *Boolean) UnmarshalJSON(bytes []byte) error {
	boolValue, err := strconv.ParseBool(string(bytes))
	*b = Boolean(boolValue)
	return err
}

// GormDataType gorm 定义数据库字段类型
// Author Kevin·CC
func (b *Boolean) GormDataType() string {
	return "tinyint"
}
