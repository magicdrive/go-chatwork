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
	"github.com/magicdrive/go-chatwork/optional"
	"github.com/stretchr/testify/assert"
)

func TestGetRoomList(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	[
	  {
		"room_id": 123,
		"name": "Group Chat Name",
		"type": "group",
		"role": "admin",
		"sticky": false,
		"unread_num": 10,
		"mention_num": 1,
		"mytask_num": 0,
		"message_num": 122,
		"file_num": 10,
		"task_num": 17,
		"icon_path": "https://example.com/ico_group.png",
		"last_update_time": 1298905200
	  }
	]
	`

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s%s", api.ApiEndpoint, target.ResourceName),
		httpmock.NewStringResponder(http.StatusOK, mock_json),
	)

	actual, err := target.List()
	assert.Nil(t, err)

	expected := make([]resource.RoomData, 0, 32)
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestCreateRooms(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "room_id": 1234
	}
	`

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s%s", api.ApiEndpoint, target.ResourceName),
		httpmock.NewStringResponder(http.StatusOK, mock_json),
	)

	params := resource.RoomsCreateParam{
		Description:        optional.String("group chat description"),
		IconPreset:         resource.IconPresetMeeting,
		Link:               optional.NilBool(),
		LinkCode:           optional.NilString(),
		LinkNeedAcceptance: optional.NilBool(),
		MembersAdminIds:    []int{123, 542, 1001},
		MembersMemberIds:   optional.IntArray(21, 344),
		MembersReadonlyIds: optional.IntArray(15, 133),
		Name:               "Website renewal project",
	}

	actual, err := target.Create(params)
	assert.Nil(t, err)

	expected := resource.RoomCreateData{}
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestGetRoom(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "room_id": 123,
	  "name": "Group Chat Name",
	  "type": "group",
	  "role": "admin",
	  "sticky": false,
	  "unread_num": 10,
	  "mention_num": 1,
	  "mytask_num": 0,
	  "message_num": 122,
	  "file_num": 10,
	  "task_num": 17,
	  "icon_path": "https://example.com/ico_group.png",
	  "last_update_time": 1298905200
	}
	`

	httpmock.RegisterResponder(http.MethodGet,
		fmt.Sprintf("%s%s/%d", api.ApiEndpoint, target.ResourceName, room_id),
		httpmock.NewStringResponder(http.StatusOK, mock_json),
	)

	actual, err := target.Get(room_id)
	assert.Nil(t, err)

	expected := resource.RoomData{}
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestUpdateRoom(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "room_id": 1234
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s/%d", api.ApiEndpoint, target.ResourceName, room_id),
		httpmock.NewStringResponder(http.StatusOK, mock_json),
	)

	params := resource.RoomsUpdateParam{
		Description: optional.String("updated description."),
		IconPreset:  resource.IconPresetIdea,
		Name:        optional.String("Updated Name"),
	}

	actual, err := target.Update(room_id, params)
	assert.Nil(t, err)

	expected := resource.RoomUpdateData{}
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestDeleteRoom(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "room_id": 1234
	}
	`

	httpmock.RegisterResponder(http.MethodDelete,
		fmt.Sprintf("%s%s/%d", api.ApiEndpoint, target.ResourceName, room_id),
		httpmock.NewStringResponder(http.StatusOK, mock_json),
	)

	err := target.Delete(room_id, resource.RoomLeave)
	assert.Nil(t, err)

}
