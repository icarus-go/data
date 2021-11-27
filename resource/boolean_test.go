package resource

import (
	"database/sql"
	"testing"
)

func Test_Boolean_SQLNullBoolean(t *testing.T) {
	nullBool := sql.NullBool{}

	if err := nullBool.Scan(3); err != nil {
		t.Fatal(err)
		return
	}

	t.Log(nullBool.Bool)
}
