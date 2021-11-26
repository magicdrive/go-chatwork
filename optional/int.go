package optional

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/goccy/go-json"
)

type NullableInt struct {
	Value  int64
	AsNull bool
}

func Int(v int64, as_null bool) NullableInt {
	return NullableInt{
		Value:  v,
		AsNull: as_null,
	}
}

func (c NullableInt) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		c.AsNull = true
		return nil
	}

	if err := json.Unmarshal(data, &c.Value); err != nil {

		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return fmt.Errorf("Couldn't unmarshal number string: %v", err)
		}

		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return fmt.Errorf("A value that cannot be interpreted as a int: %v", err)
		}
		c.Value = n
		c.AsNull = false
		return nil
	}

	c.AsNull = true
	return nil

}

func (c NullableInt) MarshalJSON() ([]byte, error) {
	if c.AsNull {
		return []byte("nil"), nil
	} else {
		return []byte(strconv.FormatInt(c.Value, 10)), nil
	}
}
