package param

import (
	"fmt"
	"strings"
)

// StringArrayParam chatwork api require string array param type
type StringArrayParam struct {
	Values []string
}

// StringArray new chatwork api require string array
func StringArray(a ...string) *StringArrayParam {
	result := &StringArrayParam{
		Values: a,
	}
	return result
}

// MarshalJSON StringArrayParam json marshaler interface.
func (c *StringArrayParam) MarshalJSON() ([]byte, error) {
	size := len(c.Values)
	if size <= 0 {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, strings.Join(c.Values, ","))), nil
}
