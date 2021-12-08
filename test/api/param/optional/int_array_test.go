package optional_test

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/stretchr/testify/assert"
)

type IntArrayTest struct {
	Ids *optional.NullableIntArray
}

func TestOptionalParamterIntArrayPresentMarshal(t *testing.T) {

	p := &IntArrayTest{Ids: optional.IntArray(1, 2, 3)}

	b, err := json.Marshal(p)

	assert.Nil(t, err)

	actual := string(b)

	expected := `{"Ids":"1,2,3"}`

	assert.Equal(t, expected, actual)
}

func TestOptionalParamterIntArrayEmptyMarshal(t *testing.T) {

	p := &IntArrayTest{Ids: optional.EmptyIntArray()}

	b, err := json.Marshal(p)

	assert.Nil(t, err)

	actual := string(b)

	expected := `{"Ids":null}`

	assert.Equal(t, expected, actual)
}
