package resource

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"pmo-test4.yz-intelligence.com/kit/data/resource/snowflake"
	"strconv"
	"strings"
)

//SnowFlakeID 雪花算法ID
//  Author: Kevin·CC
//  Description: 避免后端人员疏忽忘记书写jsontag中的string导致精度丢失的问题,直接重写unmarshal以及marshal方法强行定制规范
type SnowFlakeID uint64

//NewID
//  Author: Kevin·CC
//  Description: 创建雪花算法ID
//  Return *SnowFlakeID 雪花算法ID
//  Return error 相同机器上可能会出现ID重复的问题(可能性较低,能明确架构时可以忽略错误)
func NewID() (SnowFlakeID, error) {
	id, err := snowflake.NewID()
	if err != nil {
		return 0, err
	}
	flakeID := SnowFlakeID(id)
	return flakeID, err
}

// String
//  Description: 字符串值
//  Author: Kevin·CC
//  return string
func (s SnowFlakeID) String() string {
	return strconv.FormatUint(s.Uint64(), 10)
}

//Uint64
//  Author: Kevin·CC
//  Description: 转Uint64
//  Return uint64
func (s SnowFlakeID) Uint64() uint64 {
	return uint64(s)
}

// Scan 扫描
// Author Kevin·CC
func (s *SnowFlakeID) Scan(value interface{}) error {
	nullInt64 := sql.NullInt64{}
	err := nullInt64.Scan(value)
	*s = SnowFlakeID(nullInt64.Int64)
	return err
}

// Value 值
// Author Kevin·CC
func (s SnowFlakeID) Value() (driver.Value, error) {
	u := s.Uint64()
	return driver.Value(u), nil
}

// MarshalJSON 序列化
// Author Kevin·CC
func (s *SnowFlakeID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%v"`, *s)), nil
}

// UnmarshalJSON 反序列化
// Author Kevin·CC
func (s *SnowFlakeID) UnmarshalJSON(bytes []byte) error {
	val := string(bytes)

	if val = strings.TrimSpace(val); val == "" {
		return nil
	}

	val = strings.ReplaceAll(val, "\"", "")

	snowFlakeID, err := strconv.ParseUint(val, 10, 64)
	*s = SnowFlakeID(snowFlakeID)
	return err
}

// GormDataType gorm 定义数据库字段类型
// Author Kevin·CC
func (s *SnowFlakeID) GormDataType() string {
	return "bigint"
}
