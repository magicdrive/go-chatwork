package rooms_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/magicdrive/go-chatwork/api/resource/rooms"
	"github.com/stretchr/testify/assert"
)

func TestListRoomsTasks(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1

	target := client.Rooms().Tasks()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	[
	  {
		"task_id": 3,
		"account": {
		  "account_id": 123,
		  "name": "Bob",
		  "avatar_image_url": "https://example.com/abc.png"
		},
		"assigned_by_account": {
		  "account_id": 456,
		  "name": "Anna",
		  "avatar_image_url": "https://example.com/def.png"
		},
		"message_id": "13",
		"body": "buy milk",
		"limit_time": 1384354799,
		"status": "open",
		"limit_type": "date"
	  }
	]
	`

	httpmock.RegisterResponder(http.MethodGet,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	params := chatwork.RoomTasksListParam{
		AccountID:           optional.Int(101),
		AssignedByAccountID: optional.Int(78),
		Status:              chatwork.RoomTaskBodyStatusDone,
	}

	actual, err := target.List(roomID, params)
	assert.Nil(t, err)

	expected := make([]rooms.RoomTaskData, 0, 32)
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestCreateRoomsTask(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1

	target := client.Rooms().Tasks()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "task_ids": [123,124]
	}
	`

	httpmock.RegisterResponder(http.MethodPost,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	params := chatwork.RoomTaskPostParam{
		Body:      "Buy milk",
		Limit:     optional.Int64(1385996399),
		LimitType: chatwork.RoomTaskLimitTypeDate,
		ToIDs:     param.IntArray(1, 3, 6),
	}

	actual, err := target.Create(roomID, params)
	assert.Nil(t, err)

	expected := rooms.RoomTaskPostData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestGetRoomsTask(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1
	taskID := 1

	target := client.Rooms().Tasks()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "task_id": 3,
	  "account": {
		"account_id": 123,
		"name": "Bob",
		"avatar_image_url": "https://example.com/abc.png"
	  },
	  "assigned_by_account": {
		"account_id": 456,
		"name": "Anna",
		"avatar_image_url": "https://example.com/def.png"
	  },
	  "message_id": "13",
	  "body": "buy milk",
	  "limit_time": 1384354799,
	  "status": "open",
	  "limit_type": "date"
	}
	`

	httpmock.RegisterResponder(http.MethodGet,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName+"/%d", roomID, taskID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Get(roomID, taskID)
	assert.Nil(t, err)

	expected := rooms.RoomTaskData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestUpdateRoomsTask(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1
	taskID := 1

	target := client.Rooms().Tasks()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "task_id": 1234
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName+"/%d/status", roomID, taskID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	actual, err := target.Update(roomID, taskID, chatwork.RoomTaskBodyStatusDone)
	assert.Nil(t, err)

	expected := rooms.RoomTaskPostData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
