package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/resource/my"
	"github.com/magicdrive/go-chatwork/optional"
	"github.com/stretchr/testify/assert"
)

func TestGetMyStatus(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.My().Status()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "unread_room_num": 2,
	  "mention_room_num": 1,
	  "mytask_room_num": 3,
	  "unread_num": 12,
	  "mention_num": 1,
	  "mytask_num": 8
	}
	`

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s", api.ApiEndpoint, target.ResourceName),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Get()
	assert.Nil(t, err)

	expected := my.StatusData{}
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestMyTasks(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.My().Tasks()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	[
	  {
		"task_id": 3,
		"room": {
		  "room_id": 5,
		  "name": "Group Chat Name",
		  "icon_path": "https://example.com/ico_group.png"
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

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s", api.ApiEndpoint, target.ResourceName),
		httpmock.NewStringResponder(200, mock_json),
	)

	params := my.TasksListParam{
		AssignedByAccountId: optional.Int(456),
		Status:              my.TaskStatusOpen,
	}

	actual, err := target.List(params)
	assert.Nil(t, err)

	expected := make([]my.TaskData, 0, 32)
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
