package resource

// OptimisticLock 乐观锁
//type OptimisticLock struct {
//	Current uint64 `json:"current"`
//	Before  uint64 `json:"-"`
//}
//
//// NewOptimisticLock
////  Author: Kevin·CC
////  Description: 创建一个乐观锁
////  Return *OptimisticLock 乐观锁
//func NewOptimisticLock() *OptimisticLock {
//	id, _ := NewID() // 生成一个雪花ID作为乐观锁
//
//	return &OptimisticLock{
//		Current: id.Uint64(),
//	}
//}
//
//// Scan 扫描
////  Author Kevin·CC
//func (o *OptimisticLock) Scan(value interface{}) error {
//	nullInt64 := sql.NullInt64{}
//	err := nullInt64.Scan(value)
//	o.Current = uint64(nullInt64.Int64)
//	return err
//}
//
//// Value 值
////  Author Kevin·CC
//func (o OptimisticLock) Value() (driver.Value, error) {
//	return driver.Value(o.Current), nil
//}
//
//// MarshalJSON 序列化
////  Author Kevin·CC
//func (o *OptimisticLock) MarshalJSON() ([]byte, error) {
//	return []byte(fmt.Sprintf(`"%v"`, *o)), nil
//}
//
//// UnmarshalJSON 反序列化
////  Author Kevin·CC
//func (o *OptimisticLock) UnmarshalJSON(bytes []byte) error {
//	val := string(bytes)
//
//	if val = strings.TrimSpace(val); val == "" {
//		return nil
//	}
//
//	val = strings.ReplaceAll(val, "\"", "")
//
//	snowFlakeID, err := strconv.ParseUint(val, 10, 64)
//	if err != nil {
//		return err
//	}
//	o.Before = snowFlakeID
//	o.Current = o.Before
//	return err
//}
//
//// GormDataType gorm 定义数据库字段类型
////  Author Kevin·CC
//func (o *OptimisticLock) GormDataType() string {
//	return "bigint"
//}
//
////BeforeCreate
////  Author: Kevin·CC
////  Description: 创建前设置值
////  Param tx 事务
////  Return error 错误信息
//func (o *OptimisticLock) BeforeCreate(*gorm.DB) error {
//	id, _ := NewID()
//	*o = OptimisticLock(id)
//	return nil
//}
//
//func (o *OptimisticLock) Scopes() func(tx *gorm.DB) *gorm.DB {
//	return func(tx *gorm.DB) *gorm.DB {
//		tx.Where("")
//	}
//}
