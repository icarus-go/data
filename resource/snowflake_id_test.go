package resource

import (
	"github.com/icarus-go/data/json"
	"testing"
)

type Temp struct {
	ID       SnowFlakeID `json:"id"`
	UserName string      `json:"userName"`
}

func TestSnowFlakeIDUnmarshal(t *testing.T) {
	//id, err := NewSnowFlakeID()
	//if err != nil {
	//	return
	//}
	//marshal, err := json.Marshal(&Temp{
	//	ID:       *id,
	//	UserName: "",
	//})
	//if err != nil {
	//	return
	//}
	//
	//print(string(marshal))
}

func TestSnowFlakeID_Marshal(t *testing.T) {

	temp := Temp{}

	content := `{"id":"382322321924292972","userName":"fuck"}`

	err := json.Unmarshal([]byte(content), &temp)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	print(temp.ID.String())
}
