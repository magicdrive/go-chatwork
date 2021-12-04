package param

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
)

type IntArrayParam struct{ Values []int }

func IntArray(a ...int) *IntArrayParam {
	result := &IntArrayParam{
		Values: a,
	}
	return result
}

func (c *IntArrayParam) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		c.Values = []int{}
		return nil
	}

	if err := json.Unmarshal(data, &c.Values); err != nil {
		t := strings.Split(string(data), ",")
		var values = make([]int, 0, len(t))

		for _, item := range t {
			node, err := strconv.Atoi(item)
			if err != nil {
				return fmt.Errorf("A value that cannot be interpreted as a int: %v", err)
			}
			values = append(values, node)
		}
		c.Values = values

		return nil
	}

	return nil

}

func (c *IntArrayParam) MarshalJSON() ([]byte, error) {
	size := len(c.Values)
	if size <= 0 {
		return []byte("null"), nil
	} else {
		return []byte(strings.Trim(strings.Replace(fmt.Sprint(c.Values), " ", ",", -1), "[]")), nil
	}
}
