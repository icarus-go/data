package resource

import (
	"fmt"
	"github.com/icarus-go/data/json"
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
	idCard := "440508199606192918"
	println(string(idCard[17]))
}

func Test_has_value_Union(t *testing.T) {
	strings := NewStrings("1", "2", "3")

	union := strings.Union("3", "4", "5", "6")

	fmt.Printf("%v", union.Array())
}

func Test_not_value_Union(t *testing.T) {
	strings := NewStrings()

	union := strings.Union("3", "4", "5", "6")

	fmt.Printf("%v", union.Array())
}

func Test_has_value_intersection(t *testing.T) {
	strings := NewStrings([]string{"1", "2", "3", "4"}...)

	intersection := strings.Intersection("1", "4", "5")
	fmt.Printf("%v", intersection.Array())
}

func Test_not_value_intersection(t *testing.T) {
	strings := NewStrings()

	intersection := strings.Intersection("1", "4", "5")

	fmt.Printf("%v", intersection.Array())
}

func Test_has_value_difference(t *testing.T) {
	strings := NewStrings("1", "2", "3", "4")

	difference := strings.DifferenceSet("5", "4", "3", "1")

	fmt.Printf("%v", difference.Array())
}
