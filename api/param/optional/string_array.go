package optional

import (
	"github.com/magicdrive/go-chatwork/api/param"
)

type NullableStringArray struct {
	values *param.StringArrayParam
	asNull bool
}

func StringArray(a ...string) *NullableStringArray {
	result := &NullableStringArray{
		values: param.StringArray(a...),
		asNull: false,
	}
	return result
}

func EmptyStringArray() *NullableStringArray {
	result := &NullableStringArray{
		values: &param.StringArrayParam{},
		asNull: true,
	}
	return result
}

func (c *NullableStringArray) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	} else {
		b, err :=  c.values.MarshalJSON()
		return b, err
	}
}
