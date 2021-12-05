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

func TestListRoomsMessages(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1

	target := client.Rooms().Message()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	[
	  {
		"account_id": 123,
		"role": "member",
		"name": "John Smith",
		"chatwork_id": "tarochatworkid",
		"organization_id": 101,
		"organization_name": "Hello Company",
		"department": "Marketing",
		"avatar_image_url": "https://example.com/abc.png"
	  }
	]
	`

	httpmock.RegisterResponder(http.MethodGet,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName, room_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	force := optional.NilBool()

	actual, err := target.List(room_id, force)
	assert.Nil(t, err)

	expected := make([]rooms.RoomMessageData, 0, 32)
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestPostRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	room_id := 1

	target := client.Rooms().Message()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "message_id": "1234"
	}
	`

	httpmock.RegisterResponder(http.MethodPost,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName, room_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	params := chatwork.RoomMessagePostParam{
		Body:       "Hello.",
		SelfUnread: optional.BoolFalse(),
	}

	actual, err := target.Post(room_id, params)
	assert.Nil(t, err)

	expected := rooms.RoomMessagePostData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestReadRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1
	message_id := optional.String("1")

	target := client.Rooms().Message()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "unread_num": 461,
	  "mention_num": 0
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName+`/read`, room_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Read(room_id, message_id)
	assert.Nil(t, err)

	expected := rooms.RoomMessageReadStatusData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestUnreadRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1
	message_id := optional.String("1")

	target := client.Rooms().Message()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "unread_num":  3,
	  "mention_num": 0
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName+`/unread`, room_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Unread(room_id, message_id)
	assert.Nil(t, err)

	expected := rooms.RoomMessageReadStatusData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestGetRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1
	message_id := "1"

	target := client.Rooms().Message()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "message_id": "5",
	  "account": {
		"account_id": 123,
		"name": "Bob",
		"avatar_image_url": "https://example.com/ico_avatar.png"
	  },
	  "body": "Hello Chatwork!",
	  "send_time": 1384242850,
	  "update_time": 0
	}
	`

	httpmock.RegisterResponder(http.MethodGet,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName+"/%s", room_id, message_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Get(room_id, message_id)
	assert.Nil(t, err)

	expected := rooms.RoomMessageData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestEditRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1
	message_id := "1"

	target := client.Rooms().Message()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "message_id": "1234"
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName+"/%s", room_id, message_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	body := "Hello. Chatwork!"

	actual, err := target.Edit(room_id, message_id, body)
	assert.Nil(t, err)

	expected := rooms.RoomMessagePostData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestDeleteRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1
	message_id := "1"

	target := client.Rooms().Message()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "message_id": "5",
	  "account": {
		"account_id": 123,
		"name": "Bob",
		"avatar_image_url": "https://example.com/ico_avatar.png"
	  },
	  "body": "Hello Chatwork!",
	  "send_time": 1384242850,
	  "update_time": 0
	}
	`

	httpmock.RegisterResponder(http.MethodDelete,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName+"/%s", room_id, message_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Delete(room_id, message_id)
	assert.Nil(t, err)

	expected := rooms.RoomMessagePostData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
