package optional

import (
	"errors"
	"strconv"
)

// NullableInt chatwork api optional int param (with nullablity)
type NullableInt struct {
	value  int
	asNull bool
	valid  bool
}

// Int new chatwork api optional int
func Int(v int) *NullableInt {
	return &NullableInt{
		value:  v,
		asNull: false,
		valid:  true,
	}
}

// NilInt new chatwork api optional nil int
func NilInt() *NullableInt {
	return &NullableInt{
		value:  0,
		asNull: true,
		valid:  false,
	}
}

// NewNullableInt new chatwork api optional int with detailed.
func NewNullableInt(v int, as_null bool) *NullableInt {
	return &NullableInt{
		value:  v,
		asNull: as_null,
		valid:  false,
	}
}

// Valid mark as validated to this struct.
func (c *NullableInt) Valid() *NullableInt {
	c.valid = true
	return c
}

// Invalid mark as invalid to this struct.
func (c *NullableInt) Invalid() *NullableInt {
	c.valid = false
	return c
}

// Get get validated value in this struct.
func (c *NullableInt) Get() int {
	if c.valid {
		return c.value
	}
	panic(errors.New("NullableInt: `Get` was called without being validated.(*NullableInt.Valid())"))
}

// Value get value in this struct.
func (c *NullableInt) Value() (int, error) {
	if c.asNull {
		return 0, errors.New("This value is null.")
	}
	return c.value, nil
}

// ToNullableString convert to *NullableString
func (c *NullableInt) ToNullableString() *NullableString {
	if c.asNull {
		return NilString()
	}

	s := strconv.Itoa(c.value)

	return String(s)
}

// IsNull return asNull status
func (c *NullableInt) IsNull() bool {
	return c.asNull
}

// IsPresent return revert asNull status
func (c *NullableInt) IsPresent() bool {
	return !c.asNull
}

// MarshalJSON NullableInt json marshaler interface.
func (c *NullableInt) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	}
	return []byte(strconv.Itoa(c.value)), nil
}
