package optional

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/goccy/go-json"
)

type NullableString struct {
	value  string
	asNull bool
	valid  bool
}

func String(v string) *NullableString {
	return &NullableString{
		value:  v,
		asNull: false,
		valid:  true,
	}
}

func NilString() *NullableString {
	return &NullableString{
		value:  "",
		asNull: true,
		valid:  false,
	}
}

func NewNullableString(v string, as_null bool) *NullableString {
	return &NullableString{
		value:  v,
		asNull: as_null,
		valid:  false,
	}
}

func (c *NullableString) Valid() *NullableString {
	c.valid = true
	return c
}

func (c *NullableString) Invalid() *NullableString {
	c.valid = false
	return c
}

func (c *NullableString) Get() string {
	if c.valid {
		return c.value
	}
	panic(errors.New("NullableString: `Get` was called without being validated.(*NullableString.Valid())"))
}

func (c *NullableString) Value() (string, error) {
	if c.asNull {
		return "", errors.New("NullableString: This value is null.")
	}
	return c.value, nil
}

func (c *NullableString) IsNull() bool {
	return c.asNull
}

func (c *NullableString) IsPresent() bool {
	return !c.asNull
}

func (c *NullableString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		c.asNull = true
		return nil
	}

	if err := json.Unmarshal(data, &c.value); err != nil {
		return err
	}

	c.asNull = true
	return nil

}

func (c *NullableString) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	} else {
		return []byte(fmt.Sprintf(`"%v"`, c.value)), nil
	}
}
