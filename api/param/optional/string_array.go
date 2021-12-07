package optional

import (
	"github.com/magicdrive/go-chatwork/api/param"
)

// NullableStringArray chatwork api optional string array param (with nullablity)
type NullableStringArray struct {
	values *param.StringArrayParam
	asNull bool
}

// StringArray new chatwork api optional string array
func StringArray(a ...string) *NullableStringArray {
	result := &NullableStringArray{
		values: param.StringArray(a...),
		asNull: false,
	}
	return result
}

// EmptyStringArray new chatwork api optional empty string array
func EmptyStringArray() *NullableStringArray {
	result := &NullableStringArray{
		values: &param.StringArrayParam{},
		asNull: true,
	}
	return result
}

// MarshalJSON NullableStringArray json marshaler interface.
func (c *NullableStringArray) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	}
	b, err := c.values.MarshalJSON()
	return b, err
}
