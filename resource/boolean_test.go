package resource

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func Test_Boolean_SQLNullBoolean(t *testing.T) {
	nullBool := sql.NullBool{}

	if err := nullBool.Scan(3); err != nil {
		t.Fatal(err)
		return
	}

	t.Log(nullBool.Bool)
}
func Test_Calc_days(t *testing.T) {
	date := NewDatetimeByLayout("2006-01-02 15:04:05", "2021-12-09 00:08:56")

	sub := int(time.Now().Sub(date.Time).Hours()) // 相差多少个小时

	days := sub / 24 // 多少天

	if (sub % 24) > 0 {
		days = days + 1
	}

	fmt.Println(days)
}
