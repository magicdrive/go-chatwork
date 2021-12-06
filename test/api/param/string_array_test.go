package param

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/stretchr/testify/assert"
)

type StringArrayTest struct {
	Names *param.StringArrayParam
}

func TestParamterStringPresentMarshal(t *testing.T) {

	p := &StringArrayTest{Names: param.StringArray("foo", "bar")}

	b, err := json.Marshal(p)

	assert.Nil(t, err)

	actual := string(b)

	expected := `{"Names":"foo,bar"}`

	assert.Equal(t, expected, actual)
}

