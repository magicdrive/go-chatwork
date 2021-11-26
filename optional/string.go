package optional

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/goccy/go-json"
)

type nullableString struct {
	value  string
	asNull bool
}

func String(v string, as_null bool) *nullableString {
	return &nullableString{
		value:  v,
		asNull: as_null,
	}
}

func (c *nullableString) Value() (string, error) {
	if c.asNull {
		return "", errors.New("this value is null.")
	}
	return c.value, nil
}

func (c *nullableString) IsNull() bool {
	return c.asNull

}

func (c *nullableString) UnmarshalJSON(data []byte) error {
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

func (c *nullableString) MarshalJSON() ([]byte, error) {
	if c.asNull {
		return []byte("null"), nil
	} else {
		return []byte(fmt.Sprintf(`"%v"`, c.value)), nil
	}
}
