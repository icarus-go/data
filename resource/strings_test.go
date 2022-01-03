package resource

import (
	"pmo-test4.yz-intelligence.com/kit/data/json"
	"testing"
)

type TempString struct {
	Values Strings `json:"value"`
	ID     uint64  `json:"id"`
}

func TestStrings_Marshal(t *testing.T) {
	tempString := TempString{
		Values: *NewStrings("a", "b", "c"),
		ID:     1,
	}

	marshal, err := json.Marshal(tempString)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	print(string(marshal))
}

func TestStrings_Unmarshal(t *testing.T) {

	tempString := TempString{}
	content := `{"value":["1","2"],"id":1}`
	err := json.Unmarshal([]byte(content), &tempString)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	print(tempString.Values.String())
}

func TestStrings_In(t *testing.T) {

	strings := NewStrings("1", "2", "3")

	if strings.In("1") {
		print("包含1")
	}

	if strings.In("1", "3") {
		print("包含1,3")
	}

	if strings.In("1", "3") {
		print("包含1,3")
	}

}

func TestStrings_OrIn(t *testing.T) {

	strings := NewStrings("1", "2", "3")

	if hit, ok := strings.OrIn("1", "4"); ok {
		print("包含", hit)
	}

}

func Test_goto(t *testing.T) {
	//var c := [1,2,3,4,5,6]
	//for i:= 0;
}
