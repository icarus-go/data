package snowflake

import (
	"testing"
)

func TestSnowFlake(t *testing.T) {

	id, err := NewID()

	if err != nil {
		t.Error("生成ID失败：", err)
	}

	if id == 0 {
		t.Error("生成ID失败：", err)
	}

}
