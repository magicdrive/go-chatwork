package optional

import (
	"errors"
)

// NullableBool chatwork api optional bool param (with nullablity)
type NullableBool struct {
	value  bool
	asNull bool
	valid  bool
}

// Bool new chatwork api optional bool
func Bool(v bool) *NullableBool {
	return &NullableBool{
		value:  v,
		asNull: false,
		valid:  true,
	}
}

// BoolTrue new chatwork api optional true bool
func BoolTrue() *NullableBool {
	return &NullableBool{
		value:  true,
		asNull: false,
		valid:  true,
	}
}

// BoolFalse new chatwork api optional false bool
func BoolFalse() *NullableBool {
	return &NullableBool{
		value:  false,
		asNull: false,
		valid:  true,
	}
}

// NilBool new chatwork api optional nil bool
func NilBool() *NullableBool {
	return &NullableBool{
		value:  false,
		asNull: true,
		valid:  false,
	}
}

// NewNullableBool new chatwork api optional bool with detailed.
func NewNullableBool(v bool, asNull bool) *NullableBool {
	return &NullableBool{
		value:  v,
		asNull: asNull,
		valid:  false,
	}
}

// Valid mark as validated to this struct.
func (c *NullableBool) Valid() *NullableBool {
	c.valid = true
	return c
}

// Invalid mark as invalid to this struct.
func (c *NullableBool) Invalid() *NullableBool {
	c.valid = false
	return c
}

// Get get validated value in this struct.
func (c *NullableBool) Get() bool {
	if c.valid {
		return c.value
	}
	panic(errors.New("NullableBool: `Get` was called without being validated.(*NullableBool.Valid())"))
}

// Value get value in this struct.
func (c *NullableBool) Value() (bool, error) {
	if c.asNull {
		return false, errors.New("This value is null.")
	}
	return c.value, nil
}

// IsNull return revert asNull status
func (c *NullableBool) IsNull() bool {
	return c.asNull
}

// IsPresent return revert asNull status
func (c *NullableBool) IsPresent() bool {
	return !c.asNull
}

// ToNullableString convert to *NullableString
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

// MarshalJSON NullableBool json marshaler interface.
func (c *NullableBool) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	}
	if c.value {
		return []byte("1"), nil
	} else {
		return []byte("0"), nil
	}
}
