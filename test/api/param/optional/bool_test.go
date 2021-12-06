package optional

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/stretchr/testify/assert"
)

type BoolTest struct {
	Status *optional.NullableBool
}

func TestOptionalParamterBoolTruePresentMarshal(t *testing.T) {

	p := &BoolTest{Status: optional.BoolTrue()}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"Status":1}`

	assert.Equal(t, expected, actual)
}

func TestOptionalParamterBoolFalsePresentMarshal(t *testing.T) {

	p := &BoolTest{Status: optional.BoolFalse()}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"Status":0}`

	assert.Equal(t, expected, actual)
}

func TestOptionalParamterBoolEmptyMarshal(t *testing.T) {

	p := &BoolTest{Status: optional.NilBool()}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"Status":null}`

	assert.Equal(t, expected, actual)
}
