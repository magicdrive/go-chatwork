package optional

import (
	"errors"
	"strconv"
)

type NullableInt struct {
	value  int
	asNull bool
	valid  bool
}

func Int(v int) *NullableInt {
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

func NewNullableInt(v int, as_null bool) *NullableInt {
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

func (c *NullableInt) Get() int {
	if c.valid {
		return c.value
	}
	panic(errors.New("NullableInt: `Get` was called without being validated.(*NullableInt.Valid())"))
}

func (c *NullableInt) Value() (int, error) {
	if c.asNull {
		return 0, errors.New("this value is null.")
	}
	return c.value, nil
}

func (c *NullableInt) ToNullableString() *NullableString {
	if c.asNull {
		return NilString()
	}

	s := strconv.Itoa(c.value)

	return String(s)
}

func (c *NullableInt) IsNull() bool {
	return c.asNull
}

func (c *NullableInt) IsPresent() bool {
	return !c.asNull
}

func (c *NullableInt) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	} else {
		return []byte(strconv.Itoa(c.value)), nil
	}
}
