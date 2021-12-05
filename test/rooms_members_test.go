package test

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

func TestListRoomsMembers(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	room_id := 1

	target := client.Rooms().Members()

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

	expected := make([]rooms.RoomMemberData, 0, 32)
	err = json.Unmarshal([]byte(mock_json), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestUpdateRoomsMembers(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	room_id := 1

	target := client.Rooms().Members()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mock_json := `
	{
	  "admin": [123, 542, 1001],
	  "member": [10, 103],
	  "readonly": [6, 11]
	}
	`

	httpmock.RegisterResponder(http.MethodPost,
		fmt.Sprintf("%s%s", client.Client.ApiEndpoint, fmt.Sprintf(target.ResourceName, room_id)),
		httpmock.NewStringResponder(200, mock_json),
	)

	params := chatwork.RoomMembersUpdateParam{
		MembersAdminIds:    param.IntArray(1, 2, 3),
		MembersMemberIds:   optional.IntArray(10, 11, 12, 13),
		MembersReadonlyIds: optional.IntArray(111, 112, 114),
	}

	actual, err := target.Update(room_id, params)
	assert.Nil(t, err)

	expected := rooms.RoomMembersAuthorityData{}
	err = json.Unmarshal([]byte(mock_json), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
