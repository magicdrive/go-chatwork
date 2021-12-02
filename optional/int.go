package optional

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/goccy/go-json"
)

type NullableInt struct {
	value  int64
	asNull bool
	valid  bool
}

func IntArray(a ...int64) []*NullableInt {
	result := make([]*NullableInt, 0, 32)
	for _, v := range a {
		item := &NullableInt{
			value:  v,
			asNull: false,
			valid:  true,
		}
		result = append(result, item)
	}
	return result
}

func IntEmptyArray(a ...int64) []*NullableInt {
	return []*NullableInt{}
}

func Int(v int64) *NullableInt {
	return &NullableInt{
		value:  v,
		asNull: false,
		valid:  true,
	}
}

func NilInt() *NullableInt {
	return &NullableInt{
		value:  0,
		asNull: true,
		valid:  false,
	}
}

func NewNullableInt(v int64, as_null bool) *NullableInt {
	return &NullableInt{
		value:  v,
		asNull: as_null,
		valid:  false,
	}
}

func (c *NullableInt) Valid() *NullableInt {
	c.valid = true
	return c
}

func (c *NullableInt) Invalid() *NullableInt {
	c.valid = false
	return c
}

func (c *NullableInt) Get() int64 {
	if c.valid {
		return c.value
	}
	panic(errors.New("NullableInt: `Get` was called without being validated.(*NullableInt.Valid())"))
}

func (c *NullableInt) Value() (int64, error) {
	if c.asNull {
		return 0, errors.New("this value is null.")
	}
	return c.value, nil
}

func (c *NullableInt) ToNullableString() *NullableString {
	if c.asNull {
		return NilString()
	}

	s := strconv.FormatInt(c.value, 10)

	return String(s)
}

func (c *NullableInt) IsNull() bool {
	return c.asNull
}

func (c *NullableInt) IsPresent() bool {
	return !c.asNull
}

func (c *NullableInt) UnmarshalJSON(data []byte) error {
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

	c.asNull = false
	return nil

}

func (c *NullableInt) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("nil"), nil
	} else {
		return []byte(strconv.FormatInt(c.value, 10)), nil
	}
}
