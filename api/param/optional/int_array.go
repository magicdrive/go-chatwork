package optional

import (
	"github.com/magicdrive/go-chatwork/api/param"
)

// NullableIntArray chatwork api optional int array param (with nullablity)
type NullableIntArray struct {
	values *param.IntArrayParam
	asNull bool
}

// IntArray new chatwork api optional int array
func IntArray(a ...int) *NullableIntArray {
	result := &NullableIntArray{
		values: param.IntArray(a...),
		asNull: false,
	}
	return result
}

// EmptyIntArray new chatwork api optional empty int array
func EmptyIntArray() *NullableIntArray {
	result := &NullableIntArray{
		values: &param.IntArrayParam{},
		asNull: true,
	}
	return result
}

// MarshalJSON NullableIntArray json marshaler interface.
func (c *NullableIntArray) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	}
	return c.values.MarshalJSON()
}
