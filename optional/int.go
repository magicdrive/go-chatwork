package optional

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/goccy/go-json"
)

type nullableInt struct {
	value  int64
	asNull bool
}

func Int(v int64, as_null bool) nullableInt {
	return nullableInt{
		value:  v,
		asNull: as_null,
	}
}

func (c *nullableInt) Value() (int64, error) {
	if c.asNull {
		return 0, errors.New("this value is null.")
	}
	return c.value, nil
}

func (c *nullableInt) IsNull() bool {
	return c.asNull

}

func (c *nullableInt) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		c.asNull = true
		return nil
	}

	if err := json.Unmarshal(data, &c.value); err != nil {

		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return fmt.Errorf("Couldn't unmarshal number string: %v", err)
		}

		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return fmt.Errorf("A value that cannot be interpreted as a int: %v", err)
		}
		c.value = n
		c.asNull = false
		return nil
	}

	c.asNull = true
	return nil

}

func (c *nullableInt) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("nil"), nil
	} else {
		return []byte(strconv.FormatInt(c.value, 10)), nil
	}
}
