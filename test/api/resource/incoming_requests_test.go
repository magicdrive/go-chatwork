package resource_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetIncomingRequests(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.IncomingRequests()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s%s", client.Client.APIEndpoint, target.ResourceName),
		httpmock.NewStringResponder(http.StatusOK, mockBody),
	)

	actual, err := target.List()
	assert.Nil(t, err)

	expected := make([]resource.IncomingRequestData, 0, 32)
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestPutIncomingRequests(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	requestID := 1

	target := client.IncomingRequests()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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
	`

	httpmock.RegisterResponder(http.MethodPut, fmt.Sprintf("%s%s/%d",
		client.Client.APIEndpoint, target.ResourceName, requestID,
	),
		httpmock.NewStringResponder(http.StatusOK, mockBody),
	)

	actual, err := target.Accept(requestID)
	assert.Nil(t, err)

	expected := resource.IncomingRequestData{}
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestDeleteIncomingRequests(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	requestID := 1

	target := client.IncomingRequests()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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
	`

	httpmock.RegisterResponder(http.MethodDelete, fmt.Sprintf("%s%s/%d",
		client.Client.APIEndpoint, target.ResourceName, requestID,
	),
		httpmock.NewStringResponder(http.StatusOK, mockBody),
	)

	err := target.Delete(requestID)
	assert.Nil(t, err)

}
