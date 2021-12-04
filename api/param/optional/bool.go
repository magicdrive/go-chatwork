package optional

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/goccy/go-json"
)

type NullableBool struct {
	value  bool
	asNull bool
	valid  bool
}

func Bool(v bool) *NullableBool {
	return &NullableBool{
		value:  v,
		asNull: false,
		valid:  true,
	}
}

func BoolTrue() *NullableBool {
	return &NullableBool{
		value:  true,
		asNull: false,
		valid:  true,
	}
}

func BoolFalse() *NullableBool {
	return &NullableBool{
		value:  false,
		asNull: false,
		valid:  true,
	}
}

func NilBool() *NullableBool {
	return &NullableBool{
		value:  false,
		asNull: true,
		valid:  false,
	}
}

func NewNullableBool(v bool, as_null bool) *NullableBool {
	return &NullableBool{
		value:  v,
		asNull: as_null,
		valid:  false,
	}
}

func (c *NullableBool) Valid() *NullableBool {
	c.valid = true
	return c
}

func (c *NullableBool) Invalid() *NullableBool {
	c.valid = false
	return c
}

func (c *NullableBool) Get() bool {
	if c.valid {
		return c.value
	}
	panic(errors.New("NullableBool: `Get` was called without being validated.(*NullableBool.Valid())"))
}

func (c *NullableBool) Value() (bool, error) {
	if c.asNull {
		return false, errors.New("this value is null.")
	}
	return c.value, nil
}

func (c *NullableBool) IsNull() bool {
	return c.asNull
}

func (c *NullableBool) IsPresent() bool {
	return !c.asNull
}

func (c *NullableBool) ToNullableString() *NullableString {
	if c.asNull {
		return NilString()
	}

	if c.value {
		return String("1")
	} else {
		return String("0")
	}
}

func (c *NullableBool) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		c.asNull = true
		return nil
	}

	if err := json.Unmarshal(data, &c.value); err != nil {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return fmt.Errorf("Couldn't unmarshal number string: %v", err)
		}

		switch str {
		case "1":
			c.value = true
		case "0":
			c.value = false
		default:
			return fmt.Errorf("A value that cannot be interpreted as a bool: %v", err)
		}
		c.asNull = false
		return nil
	}

	c.asNull = false
	return nil

}

func (c *NullableBool) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	} else {
		if c.value {
			return []byte("1"), nil
		} else {
			return []byte("0"), nil
		}
	}
}
