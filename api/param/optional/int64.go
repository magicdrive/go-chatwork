package optional

import (
	"errors"
	"strconv"
)

type NullableInt64 struct {
	value  int64
	asNull bool
	valid  bool
}

func Int64(v int64) *NullableInt64 {
	return &NullableInt64{
		value:  v,
		asNull: false,
		valid:  true,
	}
}

func NilInt64() *NullableInt64 {
	return &NullableInt64{
		value:  0,
		asNull: true,
		valid:  false,
	}
}

func NewNullableInt64(v int64, as_null bool) *NullableInt64 {
	return &NullableInt64{
		value:  v,
		asNull: as_null,
		valid:  false,
	}
}

func (c *NullableInt64) Valid() *NullableInt64 {
	c.valid = true
	return c
}

func (c *NullableInt64) Invalid() *NullableInt64 {
	c.valid = false
	return c
}

func (c *NullableInt64) Get() int64 {
	if c.valid {
		return c.value
	}
	panic(errors.New("NullableInt: `Get` was called without being validated.(*NullableInt.Valid())"))
}

func (c *NullableInt64) Value() (int64, error) {
	if c.asNull {
		return 0, errors.New("this value is null.")
	}
	return c.value, nil
}

func (c *NullableInt64) ToNullableString() *NullableString {
	if c.asNull {
		return NilString()
	}

	s := strconv.FormatInt(c.value, 10)

	return String(s)
}

func (c *NullableInt64) IsNull() bool {
	return c.asNull
}

func (c *NullableInt64) IsPresent() bool {
	return !c.asNull
}

func (c *NullableInt64) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	} else {
		return []byte(strconv.FormatInt(c.value, 10)), nil
	}
}
