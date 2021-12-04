package optional

import (
	"bytes"

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

func (c *NullableStringArray) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		c.asNull = true
		return nil
	}
	v := &param.StringArrayParam{}
	err := v.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	c.values = v
	c.asNull = false
	return nil
}

func (c *NullableStringArray) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	} else {
		return c.values.MarshalJSON()
	}
}
