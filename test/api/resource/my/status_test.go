package my

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/resource/my"
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

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s", client.Client.ApiEndpoint, target.ResourceName),
		httpmock.NewStringResponder(200, mock_json),
	)

	actual, err := target.Get()
	assert.Nil(t, err)

	expected := my.MyStatusData{}
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

