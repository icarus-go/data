package resource

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"go.uber.org/zap"
	"time"
)

const dateTimeFormat = "2006-01-02 15:04:05"

type Datetime struct {
	time.Time
}

// NewDatetime 构建 Datetime
//  Description:
//  Author: Kevin·CC
//  Param: value
//  return *Datetime
func NewDatetime(value string) *Datetime {
	_time, _ := time.ParseInLocation(dateTimeFormat, value, time.Local)
	return &Datetime{Time: _time}
}

// Now
//  Author: Kevin·CC
//  Description: 当前时间
//  Return *Datetime
func Now() *Datetime {
	return &Datetime{
		Time: time.Now(),
	}
}

// NewDatetimeByTime 通过 time.Time 构建 Datetime
//  Description:
//  Author: Kevin·CC
//  Param: t
//  return *Datetime
func NewDatetimeByTime(t time.Time) *Datetime {
	return &Datetime{Time: t}
}

// NewDatetimeByLayout 构建 Datetime by layout
//  Description:
//  Author: Kevin·CC
//  Param: layout
//  Param: value
//  return *Datetime
func NewDatetimeByLayout(layout, value string) *Datetime {
	_time, _ := time.ParseInLocation(layout, value, time.Local)
	_d := &Datetime{Time: _time}
	return _d
}

// ToDate Datetime 2 Date
//  Description:
//  Author: Kevin·CC
//  return *Date
func (d *Datetime) ToDate() *Date {
	if d == nil {
		return nil
	}
	return &Date{Time: d.Time}
}

// ToTime Datetime 2 *time.Time
// Author SliverHorn
func (d *Datetime) ToTime() *time.Time {
	if d == nil {
		return nil
	}
	return &d.Time
}

// Scan 扫描
// Author SliverHorn
func (d *Datetime) Scan(value interface{}) error {
	nullTime := &sql.NullTime{}
	if err := nullTime.Scan(value); err != nil {
		zap.L().Error("Date To Database 时间转换失败!", zap.Any("datetime", value))
		return err
	} else {
		*d = Datetime{Time: nullTime.Time}
		return nil
	}
}

// Value 值
// Author SliverHorn
func (d Datetime) Value() (driver.Value, error) {
	return driver.Value(d.Time.Format(dateTimeFormat)), nil
}

// MarshalJSON 序列化
//  Description:
//  Author: Kevin·CC
//  return []byte 字节码
//  return error 错误信息
func (d Datetime) MarshalJSON() ([]byte, error) {
	value := d.Time.Format(dateTimeFormat)
	return []byte(fmt.Sprintf(`"%s"`, value)), nil
}

// UnmarshalJSON 反序列化
//  Description:
//  Author: Kevin·CC
//  Param: b 字节码
//  return error 错误信息
func (d *Datetime) UnmarshalJSON(b []byte) error {
	if string(b) > `""` {
		_time, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), time.Local) // 为什么这里不使用上面的dateTimeFormat呢? 原因是避免字符串拼接,在大数据的情况下性能过低的问题
		if err != nil {
			zap.L().Error("时间转换失败!", zap.String("datetime", string(b)))
			return err
		}
		*d = Datetime{Time: _time}
		return nil
	}
	return nil
}

// GormDataType gorm 定义数据库字段类型
//  Description:
//  Author: Kevin·CC
//  return string datetime
func (d *Datetime) GormDataType() string {
	return "datetime"
}

// CostMonth 获取时间所属月份 2021-05-13 17:48:00 => May-21
// Author SliverHorn
func (d *Datetime) CostMonth() string {
	if d == nil {
		return ""
	}
	_time := d.Time
	a := fmt.Sprintf("%v", _time.Month())
	if len(a) > 3 { // 取月份简称
		return fmt.Sprintf("%v-%v", a[:3], _time.Year()-2000)
	}
	return fmt.Sprintf("%v-%v", a, _time.Year()-2000)
}
