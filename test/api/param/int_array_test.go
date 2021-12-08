package param_test

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/stretchr/testify/assert"
)

type IntArrayTest struct {
	Ids *param.IntArrayParam
}

func TestIntArray(t *testing.T) {
	a := []int{1, 2, 3}

	expected := &param.IntArrayParam{
		Values: a,
	}

	actual := param.IntArray(a...)

	assert.Equal(t, expected, actual)
}

func TestIntArrayParamMarshalJSON(t *testing.T) {

	p := &IntArrayTest{Ids: param.IntArray(1, 2, 3)}

	b, err := json.Marshal(p)

	assert.Nil(t, err)

	actual := string(b)

	expected := `{"Ids":"1,2,3"}`

	assert.Equal(t, expected, actual)
}
