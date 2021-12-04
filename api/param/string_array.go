package param

import (
	"fmt"
	"strings"
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

func (c *StringArrayParam) MarshalJSON() ([]byte, error) {
	size := len(c.Values)
	if size <= 0 {
		return []byte("null"), nil
	} else {
		return []byte(fmt.Sprintf(`"%s"`, strings.Join(c.Values, ","))), nil
	}
}
