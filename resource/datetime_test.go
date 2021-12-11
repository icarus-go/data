package resource

import (
	//"pmo-test4.yz-intelligence.com/kit/data/json"
	"encoding/json"
	"testing"
)

func TestResourceDatetime_Marshal(t *testing.T) {

	content := `
		{
			"datetime": "2021-12-11 21:19:49"
		}
	`

	temp := struct {
		Time Datetime `json:"datetime"`
	}{}

	_ = json.Unmarshal([]byte(content), &temp)

	println(temp.Time.Unix())

	result, _ := json.Marshal(temp)

	print(string(result))

}
