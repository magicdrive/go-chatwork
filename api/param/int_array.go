package param

import (
	"fmt"
	"strings"
)

type IntArrayParam struct {
	Values []int
}

func IntArray(a ...int) *IntArrayParam {
	result := &IntArrayParam{
		Values: a,
	}
	return result
}

func (c *IntArrayParam) MarshalJSON() ([]byte, error) {
	size := len(c.Values)
	if size <= 0 {
		return []byte("null"), nil
	} else {
		return []byte(fmt.Sprintf(`"%s"`,
			strings.Trim(strings.Replace(fmt.Sprint(c.Values), " ", ",", -1), "[]"),
		)), nil
	}
}
