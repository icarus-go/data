package constant

type Week string

const (
	// Sunday 星期日
	Sunday Week = "星期日"
	// Monday 星期一
	Monday Week = "星期一"
	// Tuesday 星期二
	Tuesday Week = "星期二"
	// Wednesday 星期三
	Wednesday Week = "星期三"
	// Thursday 星期四
	Thursday Week = "星期四"
	// Friday 星期五
	Friday Week = "星期五"
	// Saturday 星期六
	Saturday Week = "星期六"
)

// Value
//  Author: Kevin·Cai
//  Description: 值
//  Return string
func (w Week) Value() string {
	return string(w)
}

// En
//  Author: Kevin·Cai
//  Description: 英文值
//  Return string
func (w Week) En() string {
	switch w {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		return ""
	}
}
