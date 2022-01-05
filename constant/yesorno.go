package constant

import "strconv"

type YesOrNo int

const (
	// Not 否
	Not YesOrNo = 0
	// Is 是
	Is YesOrNo = 1
)

// BooleanPtr 布尔值指针
//  Description: 通用用于是或否
func (v YesOrNo) BooleanPtr() *bool {
	boolean := v.Boolean()
	return &boolean
}

// Boolean
//  Author: Kevin·Cai
//  Description: 布尔值
//  Return bool
func (v YesOrNo) Boolean() bool {
	return v == Is
}

// Value
//  Description: 值
//  Author: Kevin·Cai
//  return int
func (v YesOrNo) Value() int {
	return int(v)
}

// Ptr
//  Description: 指针值
//  Author: Kevin·Cai
//  return *int 指针值
func (v YesOrNo) Ptr() *int {
	value := v.Value()
	return &value
}

// String
//  Description: 值
//  Author: yiQiu.huang
//  return string
func (v YesOrNo) String() string {
	return strconv.Itoa(int(v))
}

// Uint
//  Author: Kevin·Cai
//  Description: Uint值
//  Return uint
func (v YesOrNo) Uint() uint {
	return uint(v)
}

// UintPtr
//  Author: Kevin·Cai
//  Description: Uint指针
//  Return *uint
func (v YesOrNo) UintPtr() *uint {
	u := v.Uint()
	return &u
}
