package optional

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/stretchr/testify/assert"
)

type IntTest struct {
	Id    *optional.NullableInt
	Count *optional.NullableInt
}

func TestOptionalParamterIntPresentMarshal(t *testing.T) {

	p := &IntTest{Id: optional.Int(11111), Count: optional.Int(2222)}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"Id":11111,"Count":2222}`

	assert.Equal(t, expected, actual)
}

func TestOptionalParamterIntEmptyMarshal(t *testing.T) {

	p := &IntTest{Id: optional.NilInt(), Count: optional.NilInt()}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"Id":null,"Count":null}`

	assert.Equal(t, expected, actual)
}
