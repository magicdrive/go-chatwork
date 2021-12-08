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

func TestListRoomsMessages(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1

	target := client.Rooms().Messages()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	force := optional.NilBool()

	actual, err := target.List(roomID, force)
	assert.Nil(t, err)

	expected := make([]rooms.RoomMessageData, 0, 32)
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestPostRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1

	target := client.Rooms().Messages()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "message_id": "1234"
	}
	`

	httpmock.RegisterResponder(http.MethodPost,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	params := chatwork.RoomMessagePostParam{
		Body:       "Hello.",
		SelfUnread: optional.BoolFalse(),
	}

	actual, err := target.Post(roomID, params)
	assert.Nil(t, err)

	expected := rooms.RoomMessagePostData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestReadRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1
	messageID := optional.String("1")

	target := client.Rooms().Messages()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "unread_num": 461,
	  "mention_num": 0
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName+`/read`, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Read(roomID, messageID)
	assert.Nil(t, err)

	expected := rooms.RoomMessageReadStatusData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestUnreadRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1
	messageID := optional.String("1")

	target := client.Rooms().Messages()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "unread_num":  3,
	  "mention_num": 0
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName+`/unread`, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Unread(roomID, messageID)
	assert.Nil(t, err)

	expected := rooms.RoomMessageReadStatusData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestGetRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1
	messageID := "1"

	target := client.Rooms().Messages()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName+"/%s", roomID, messageID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Get(roomID, messageID)
	assert.Nil(t, err)

	expected := rooms.RoomMessageData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestEditRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1
	messageID := "1"

	target := client.Rooms().Messages()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "message_id": "1234"
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName+"/%s", roomID, messageID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	body := "Hello. Chatwork!"

	actual, err := target.Edit(roomID, messageID, body)
	assert.Nil(t, err)

	expected := rooms.RoomMessagePostData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestDeleteRoomsMessage(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1
	messageID := "1"

	target := client.Rooms().Messages()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName+"/%s", roomID, messageID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Delete(roomID, messageID)
	assert.Nil(t, err)

	expected := rooms.RoomMessagePostData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
