package rooms_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/magicdrive/go-chatwork/api/resource/rooms"
	"github.com/stretchr/testify/assert"
)

func TestGetRoomsLink(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1

	target := client.Rooms().Link()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "public": true,
	  "url": "https://example.chatwork.com/g/randomcode42",
	  "need_acceptance": true,
	  "description": "Link description text"
	}
	`

	httpmock.RegisterResponder(http.MethodGet,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Get(roomID)
	assert.Nil(t, err)

	expected := rooms.RoomLinkData{}
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestCreateRoomsLink(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1

	target := client.Rooms().Link()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "file_id":3,
	  "account": {
		"account_id":123,
		"name":"Bob",
		"avatar_image_url": "https://example.com/ico_avatar.png"
	  },
	  "message_id": "22",
	  "filename": "README.md",
	  "filesize": 2232,
	  "upload_time": 1384414750
	}
	`

	httpmock.RegisterResponder(http.MethodPost,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	params := chatwork.RoomLinkParam{
		Code:           optional.String("unique-link-name"),
		Description:    optional.String("This is a public room for topic A."),
		NeedAcceptance: optional.BoolTrue(),
	}

	actual, err := target.Create(roomID, params)
	assert.Nil(t, err)

	expected := rooms.RoomLinkData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestEditRoomsLink(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1

	target := client.Rooms().Link()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "file_id":3,
	  "account": {
		"account_id":123,
		"name":"Bob",
		"avatar_image_url": "https://example.com/ico_avatar.png"
	  },
	  "message_id": "22",
	  "filename": "README.md",
	  "filesize": 2232,
	  "upload_time": 1384414750
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	params := chatwork.RoomLinkParam{
		Code:           optional.String("unique-link-name"),
		Description:    optional.String("This is a public room for topic A."),
		NeedAcceptance: optional.BoolTrue(),
	}

	actual, err := target.Edit(roomID, params)
	assert.Nil(t, err)

	expected := rooms.RoomLinkData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestDeleteRoomsLink(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1

	target := client.Rooms().Link()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "public": false
	}
	`

	httpmock.RegisterResponder(http.MethodDelete,
		fmt.Sprintf(fmt.Sprintf("%s%s", client.Client.APIEndpoint, target.ResourceName), roomID),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Delete(roomID)
	assert.Nil(t, err)

	expected := rooms.RoomLinkDeleteData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
