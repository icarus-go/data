package resource

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"pmo-test4.yz-intelligence.com/kit/data/json"
	"strings"
)

// Strings 根据元素组成的字符串
type Strings []string

//NewStrings
//  Author: Kevin·CC
//  Description: 创建字符串组,以逗号切割,拼接
//  Param val 字符串数组
//  Return *Strings 字符串组
func NewStrings(val ...string) *Strings {
	s := &Strings{}

	*s = val

	return s
}

func (s Strings) IndexOf(subString string) int {
	return strings.Index(s.String(), subString)
}

// Array
//  Author: Kevin·CC
//  Description: gorm无法识别
//  Return []string
func (s Strings) Array() []string {
	return []string(s)
}

//In
//  Author: Kevin·CC
//  Description: 是否包含元素
//  Param val 要包含的元素
//  Return bool 是否全部包含
func (s Strings) In(val ...string) bool {
	m := make(map[string]struct{}, len(s))
	for _, val := range s {
		m[val] = struct{}{}
	}

	for _, item := range val {
		if _, ok := m[item]; ok {
			continue
		}
		return false
	}
	return true
}

func (s Strings) OrIn(val ...string) (string, bool) {
	m := make(map[string]struct{}, len(s))
	for _, val := range s {
		m[val] = struct{}{}
	}

	for _, item := range val {
		if _, ok := m[item]; ok {
			return item, true
		}
	}

	return "", false
}

//NotIn
//  Author: Kevin·CC
//  Description: 是否不包含元素
//  Param val
//  Return bool
func (s Strings) NotIn(val ...string) bool {
	m := make(map[string]struct{}, len(s))
	for _, val := range s {
		m[val] = struct{}{}
	}

	for _, item := range val {
		if _, ok := m[item]; !ok {
			return true
		}
	}

	return false
}

// TODO 待支持 差集 , 合集 , 并集

// String
//  Description: 字符串值
//  Author: Kevin·CC
//  return string
func (s Strings) String() string {
	if len(s) == 0 {
		return ""
	}
	return strings.Join(s, ",")
}

func (s *Strings) trim() *Strings {
	if len(*s) < 1 {
		return s
	}

	for i := 0; i < len(*s); i++ {
		(*s)[i] = strings.TrimSpace((*s)[i])
	}
	return s
}

// Scan 扫描
// Author Kevin·CC
func (s *Strings) Scan(value interface{}) error {
	nullString := sql.NullString{}
	err := nullString.Scan(value)
	val := nullString.String
	if val != "" {
		*s = strings.Split(val, ",")
	}
	return err
}

// Value 值
// Author Kevin·CC
func (s Strings) Value() (driver.Value, error) {
	s2 := s.trim().String()
	return driver.Value(s2), nil
}

// MarshalJSON 序列化
// Author Kevin·CC
func (s *Strings) MarshalJSON() ([]byte, error) {
	_ = s.trim()
	marshal, err := json.Marshal(*s)
	return []byte(fmt.Sprintf(`%v`, string(marshal))), err
}

// UnmarshalJSON 反序列化
// Author Kevin·CC
func (s *Strings) UnmarshalJSON(bytes []byte) error {
	//val := strings.ReplaceAll(string(bytes), "\"", "")
	//val = strings.ReplaceAll(val, "[", "")
	//val = strings.ReplaceAll(val, "]", "")
	var temp []string
	err := json.Unmarshal(bytes, &temp)
	*s = temp
	return err
}

// GormDataType gorm 定义数据库字段类型
// Author Kevin·CC
func (s *Strings) GormDataType() string {
	return "varchar"
}
