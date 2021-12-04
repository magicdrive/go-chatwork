package test

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/stretchr/testify/assert"
)

type StringArrayTest struct {
	Names *optional.NullableStringArray
}

func TestOptionalParamterStringPresentMarshal(t *testing.T) {

	p := &StringArrayTest{Names: optional.StringArray("foo", "bar")}

	b, err := json.Marshal(p)

	assert.Nil(t, err)

	actual := string(b)

	expected := `{"Names":"foo,bar"}`

	assert.Equal(t, expected, actual)
}

func TestOptionalParamterStringEmptyMarshal(t *testing.T) {

	p := &StringArrayTest{Names: optional.EmptyStringArray()}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"Names":null}`

	assert.Equal(t, expected, actual)
}
