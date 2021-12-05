package test

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

	room_id := 1

	target := client.Rooms().Link()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "public": true,
	  "url": "https://example.chatwork.com/g/randomcode42",
	  "need_acceptance": true,
	  "description": "Link description text"
	}
	`

	httpmock.RegisterResponder(http.MethodGet,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName, room_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Get(room_id)
	assert.Nil(t, err)

	expected := rooms.RoomLinkData{}
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestCreateRoomsLink(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	room_id := 1

	target := client.Rooms().Link()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
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
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName, room_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	params := chatwork.RoomLinkParam{
		Code:           optional.String("unique-link-name"),
		Description:    optional.String("This is a public room for topic A."),
		NeedAcceptance: optional.BoolTrue(),
	}

	actual, err := target.Create(room_id, params)
	assert.Nil(t, err)

	expected := rooms.RoomLinkData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestEditRoomsLink(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	room_id := 1

	target := client.Rooms().Link()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
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
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName, room_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	params := chatwork.RoomLinkParam{
		Code:           optional.String("unique-link-name"),
		Description:    optional.String("This is a public room for topic A."),
		NeedAcceptance: optional.BoolTrue(),
	}

	actual, err := target.Edit(room_id, params)
	assert.Nil(t, err)

	expected := rooms.RoomLinkData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestDeleteRoomsLink(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	room_id := 1

	target := client.Rooms().Link()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "public": false
	}
	`

	httpmock.RegisterResponder(http.MethodDelete,
		fmt.Sprintf(fmt.Sprintf("%s%s", client.Client.ApiEndpoint, target.ResourceName), room_id),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Delete(room_id)
	assert.Nil(t, err)

	expected := rooms.RoomLinkDeleteData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
