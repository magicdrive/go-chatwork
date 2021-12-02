package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetContacts(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.Contacts()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	[
	  {
		"account_id": 123,
		"room_id": 322,
		"name": "John Smith",
		"chatwork_id": "tarochatworkid",
		"organization_id": 101,
		"organization_name": "Hello Company",
		"department": "Marketing",
		"avatar_image_url": "https://example.com/abc.png"
	  }
	]
	`

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s%s", api.ApiEndpoint, target.ResourceName),
		httpmock.NewStringResponder(http.StatusOK, mock_json),
	)

	actual, err := target.List()
	assert.Nil(t, err)

	expected := make([]resource.ContactData, 0, 32)
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
