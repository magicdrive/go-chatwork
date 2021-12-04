package test

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/stretchr/testify/assert"
)

type IntArrayTest struct {
	Ids *param.IntArrayParam
}

func TestParamterIntArrayPresentMarshal(t *testing.T) {

	p := &IntArrayTest{Ids: param.IntArray(1, 2, 3)}

	b, err := json.Marshal(p)

	assert.Nil(t, err)

	actual := string(b)

	expected := `{"Ids":"1,2,3"}`

	assert.Equal(t, expected, actual)
}
