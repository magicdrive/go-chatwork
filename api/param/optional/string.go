package optional

import (
	"errors"
	"fmt"
)

// NullableString chatwork api optional string param (with nullablity)
type NullableString struct {
	value  string
	asNull bool
	valid  bool
}

// String new chatwork api optional string
func String(v string) *NullableString {
	return &NullableString{
		value:  v,
		asNull: false,
		valid:  true,
	}
}

// NilString new chatwork api optional nil string
func NilString() *NullableString {
	return &NullableString{
		value:  "",
		asNull: true,
		valid:  false,
	}
}

// NewNullableString new chatwork api optional string with detailed
func NewNullableString(v string, as_null bool) *NullableString {
	return &NullableString{
		value:  v,
		asNull: as_null,
		valid:  false,
	}
}

// Valid mark validated, to this struct.
func (c *NullableString) Valid() *NullableString {
	c.valid = true
	return c
}

// Invalid mark invalid, to this struct.
func (c *NullableString) Invalid() *NullableString {
	c.valid = false
	return c
}

// Get get validated value in this struct.
func (c *NullableString) Get() string {
	if c.valid {
		return c.value
	}
	panic(errors.New("NullableString: `Get` was called without being validated.(*NullableString.Valid())"))
}

// Value get value in this struct.
func (c *NullableString) Value() (string, error) {
	if c.asNull {
		return "", errors.New(`*NullableString: This value is null.`)
	}
	return c.value, nil
}

// IsNull return asNull status
func (c *NullableString) IsNull() bool {
	return c.asNull
}

// IsPresent return reverted asNull status
func (c *NullableString) IsPresent() bool {
	return !c.asNull
}

// MarshalJSON NullableString json marshaler interface.
func (c *NullableString) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%v"`, c.value)), nil
}
