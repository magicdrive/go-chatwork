package my_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/magicdrive/go-chatwork/api/resource/my"
	"github.com/stretchr/testify/assert"
)

func TestMyTasks(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.My().Tasks()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s", client.Client.APIEndpoint, target.ResourceName),
		httpmock.NewStringResponder(200, mockBody),
	)

	params := chatwork.MyTasksListParam{
		AssignedByAccountID: optional.Int(456),
		Status:              chatwork.MyTaskStatusOpen,
	}

	actual, err := target.List(params)
	assert.Nil(t, err)

	expected := make([]my.MyTaskData, 0, 32)
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
