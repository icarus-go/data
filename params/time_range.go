package params

import "pmo-test4.yz-intelligence.com/kit/data/resource"

//DatetimeRange
//  Author: Kevin·CC
//  Description: 日期(精确到时分秒)范围
type DatetimeRange struct {
	Start *resource.Datetime `json:"start" swaggertype:"string" example:"datetime (2006-01-02 15:04:05)起始时间"` // 起始时间
	End   *resource.Datetime `json:"end" swaggertype:"string" example:"datetime (2006-01-02 15:04:05)结束时间"`   // 结束时间
}

//DateRange
//  Author: Kevin·CC
//  Description: 日期(精确到日期)范围
type DateRange struct {
	Start *resource.Date `json:"start" swaggertype:"string" example:"date (2006-01-02)起始时间"` // 起始时间
	End   *resource.Date `json:"end" swaggertype:"string" example:"date (2006-01-02)结束时间"`   // 结束时间
}
