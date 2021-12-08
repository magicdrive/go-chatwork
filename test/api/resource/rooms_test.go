package resource_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/magicdrive/go-chatwork/api/param/optional"
	"github.com/magicdrive/go-chatwork/api/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetRoomList(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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

	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("%s%s", client.Client.APIEndpoint, target.ResourceName),
		httpmock.NewStringResponder(http.StatusOK, mockBody),
	)

	actual, err := target.List()
	assert.Nil(t, err)

	expected := make([]resource.RoomData, 0, 32)
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestCreateRooms(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "room_id": 1234
	}
	`

	httpmock.RegisterResponder(http.MethodPost, fmt.Sprintf("%s%s", client.Client.APIEndpoint, target.ResourceName),
		httpmock.NewStringResponder(http.StatusOK, mockBody),
	)

	params := resource.RoomsCreateParam{
		Description:        optional.String("group chat description"),
		IconPreset:         resource.IconPresetMeeting,
		Link:               optional.NilBool(),
		LinkCode:           optional.NilString(),
		LinkNeedAcceptance: optional.NilBool(),
		MembersAdminIds:    param.IntArray(123, 542, 1001),
		MembersMemberIds:   optional.IntArray(21, 344),
		MembersReadonlyIds: optional.IntArray(15, 133),
		Name:               "Website renewal project",
	}

	actual, err := target.Create(params)
	assert.Nil(t, err)

	expected := resource.RoomsCreateData{}
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestGetRoom(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
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
		fmt.Sprintf("%s%s/%d", client.Client.APIEndpoint, target.ResourceName, roomID),
		httpmock.NewStringResponder(http.StatusOK, mockBody),
	)

	actual, err := target.Get(roomID)
	assert.Nil(t, err)

	expected := resource.RoomData{}
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestUpdateRoom(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "room_id": 1234
	}
	`

	httpmock.RegisterResponder(http.MethodPut,
		fmt.Sprintf("%s%s/%d", client.Client.APIEndpoint, target.ResourceName, roomID),
		httpmock.NewStringResponder(http.StatusOK, mockBody),
	)

	params := resource.RoomsUpdateParam{
		Description: optional.String("updated description."),
		IconPreset:  resource.IconPresetIdea,
		Name:        optional.String("Updated Name"),
	}

	actual, err := target.Update(roomID, params)
	assert.Nil(t, err)

	expected := resource.RoomUpdateData{}
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestDeleteRoom(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1

	target := client.Rooms()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "room_id": 1234
	}
	`

	httpmock.RegisterResponder(http.MethodDelete,
		fmt.Sprintf("%s%s/%d", client.Client.APIEndpoint, target.ResourceName, roomID),
		httpmock.NewStringResponder(http.StatusOK, mockBody),
	)

	err := target.Delete(roomID, resource.RoomLeave)
	assert.Nil(t, err)

}
