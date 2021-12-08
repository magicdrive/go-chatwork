package param_test

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/stretchr/testify/assert"
)

type StringArrayTest struct {
	Names *param.StringArrayParam
}

func TestStringArray(t *testing.T) {
	a := []string{"a", "b", "c"}

	expected := &param.StringArrayParam{
		Values: a,
	}

	actual := param.StringArray(a...)

	assert.Equal(t, expected, actual)
}

func TestStringArrayParamMarshalJSON(t *testing.T) {

	p := &StringArrayTest{Names: param.StringArray("foo", "bar")}

	b, err := json.Marshal(p)

	assert.Nil(t, err)

	actual := string(b)

	expected := `{"Names":"foo,bar"}`

	assert.Equal(t, expected, actual)
}
