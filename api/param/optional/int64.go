package optional

import (
	"errors"
	"strconv"
)

// NullableInt64 chatwork api optional int64 param (with nullablity)
type NullableInt64 struct {
	value  int64
	asNull bool
	valid  bool
}

// Int64 new chatwork api optional int64
func Int64(v int64) *NullableInt64 {
	return &NullableInt64{
		value:  v,
		asNull: false,
		valid:  true,
	}
}

// NilInt64 new chatwork api optional nil int64
func NilInt64() *NullableInt64 {
	return &NullableInt64{
		value:  0,
		asNull: true,
		valid:  false,
	}
}

// NewNullableInt64 new chatwork api optional int64 with detailed.
func NewNullableInt64(v int64, as_null bool) *NullableInt64 {
	return &NullableInt64{
		value:  v,
		asNull: as_null,
		valid:  false,
	}
}

// Valid mark as validated to this struct.
func (c *NullableInt64) Valid() *NullableInt64 {
	c.valid = true
	return c
}

// Invalid mark as invalid to this struct.
func (c *NullableInt64) Invalid() *NullableInt64 {
	c.valid = false
	return c
}

// Get get validated value in this struct.
func (c *NullableInt64) Get() int64 {
	if c.valid {
		return c.value
	}
	panic(errors.New("NullableInt: `Get` was called without being validated.(*NullableInt.Valid())"))
}

// Value get value in this struct.
func (c *NullableInt64) Value() (int64, error) {
	if c.asNull {
		return 0, errors.New("this value is null.")
	}
	return c.value, nil
}

// ToNullableString convert to *NullableString
func (c *NullableInt64) ToNullableString() *NullableString {
	if c.asNull {
		return NilString()
	}

	s := strconv.FormatInt(c.value, 10)

	return String(s)
}

// IsNull return asNull status
func (c *NullableInt64) IsNull() bool {
	return c.asNull
}

// IsPresent return revert asNull status
func (c *NullableInt64) IsPresent() bool {
	return !c.asNull
}

// MarshalJSON NullableInt64 json marshaler interface.
func (c *NullableInt64) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(c.value, 10)), nil
}
