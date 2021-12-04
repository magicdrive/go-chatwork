package test

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/stretchr/testify/assert"
)

type Int64Test struct {
	Id    *optional.NullableInt64
	Count *optional.NullableInt64
}

func TestOptionalParamterInt64PresentMarshal(t *testing.T) {

	p := &Int64Test{Id: optional.Int64(11111), Count: optional.Int64(2222)}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"Id":11111,"Count":2222}`

	assert.Equal(t, expected, actual)
}

func TestOptionalParamterInt64EmptyMarshal(t *testing.T) {

	p := &Int64Test{Id: optional.NilInt64(), Count: optional.NilInt64()}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"Id":null,"Count":null}`

	assert.Equal(t, expected, actual)
}
