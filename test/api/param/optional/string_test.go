package optional

import (
	"testing"

	"github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/stretchr/testify/assert"
)

type StringTest struct {
	WordPre *optional.NullableString
	WordSuf *optional.NullableString
}

func TestStringMarshalJSON(t *testing.T) {

	p := &StringTest{WordPre: optional.String("Hello"), WordSuf: optional.String("World")}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"WordPre":"Hello","WordSuf":"World"}`

	assert.Equal(t, expected, actual)
}

func TestStringMarshalJSONEmpty(t *testing.T) {

	p := &StringTest{WordPre: optional.NilString(), WordSuf: optional.NilString()}

	b, err := json.Marshal(p)

	actual := string(b)

	assert.Nil(t, err)

	expected := `{"WordPre":null,"WordSuf":null}`

	assert.Equal(t, expected, actual)
}
