package param

import (
	"fmt"
	"strings"
)

// IntArrayParam chatwork api require int array param type
type IntArrayParam struct {
	Values []int
}

// IntArray new chatwork api require int array
func IntArray(a ...int) *IntArrayParam {
	result := &IntArrayParam{
		Values: a,
	}
	return result
}

// MarshalJSON IntArrayParam json marshaler interface.
func (c *IntArrayParam) MarshalJSON() ([]byte, error) {
	size := len(c.Values)
	if size <= 0 {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`,
		strings.Trim(strings.Replace(fmt.Sprint(c.Values), " ", ",", -1), "[]"),
	)), nil
}
