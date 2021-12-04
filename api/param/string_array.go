package param

import (
	"bytes"
	"strings"

	"github.com/goccy/go-json"
)

type StringArrayParam struct {
	Values []string
}

func StringArray(a ...string) *StringArrayParam {
	result := &StringArrayParam{
		Values: a,
	}
	return result
}

func (c *StringArrayParam) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		c.Values = []string{}
		return nil
	}

	if err := json.Unmarshal(data, &c.Values); err != nil {
		c.Values = strings.Split(string(data), ",")
		return nil
	}

	return nil
}

func (c *StringArrayParam) MarshalJSON() ([]byte, error) {
	size := len(c.Values)
	if size <= 0 {
		return []byte("null"), nil
	} else {
		return []byte(strings.Join(c.Values, ",")), nil
	}
}
