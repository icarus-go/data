package constant

type Gender string

const (
	// Man 男性
	Man Gender = "男"
	// Women 女性
	Women Gender = "女"
)

func (g Gender) String() string {
	return string(g)
}
