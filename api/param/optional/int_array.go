package optional

import (
	"github.com/magicdrive/go-chatwork/api/param"
)

type NullableIntArray struct {
	values *param.IntArrayParam
	asNull bool
}

func IntArray(a ...int) *NullableIntArray {
	result := &NullableIntArray{
		values: param.IntArray(a...),
		asNull: false,
	}
	return result
}

func EmptyIntArray() *NullableIntArray {
	result := &NullableIntArray{
		values: &param.IntArrayParam{},
		asNull: true,
	}
	return result
}

func (c *NullableIntArray) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	} else {
		return c.values.MarshalJSON()
	}
}
