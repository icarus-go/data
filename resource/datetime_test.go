package resource

import (
	"fmt"
	"pmo-test4.yz-intelligence.com/kit/data/json"
	"time"

	//"encoding/json"
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

func TestResourceDatetime_MarshalByPtr(t *testing.T) {

	content := `
		{
			"datetime": "2021-12-11 21:19:49"
		}
	`

	temp := struct {
		Time *Datetime `json:"datetime"`
	}{}

	_ = json.Unmarshal([]byte(content), &temp)

	println(temp.Time.Unix())

	result, _ := json.Marshal(temp)

	print(string(result))
}

func Test_ResourceDatetime_After(t *testing.T) {
	current := time.Now()
	morning := NewDatetimeByLayout(dateTimeFormat, "2021-12-15 09:00:00")

	fmt.Printf("morning after current , current.Unix(): %d, morning.Unix(): %d , value: %v\n", current.Unix(), morning.Unix(), morning.After(current))      // 早上大于当前时间 = false
	fmt.Printf("current after morning , current.Unix(): %d, morning.Unix(): %d , value: %v\n", current.Unix(), morning.Unix(), current.After(morning.Time)) // 当前时间大于早上 = true
}
