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

func TestListRoomsMembers(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)

	roomID := 1

	target := client.Rooms().Members()

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

	expected := make([]rooms.RoomMemberData, 0, 32)
	err = json.Unmarshal([]byte(mockBody), &expected)
	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}

func TestUpdateRoomsMembers(t *testing.T) {

	client := chatwork.NewChatworkClient(`test-api-key`)
	roomID := 1

	target := client.Rooms().Members()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockBody := `
	{
	  "admin": [123, 542, 1001],
	  "member": [10, 103],
	  "readonly": [6, 11]
	}
	`

	httpmock.RegisterResponder(http.MethodPost,
		fmt.Sprintf("%s%s", client.Client.APIEndpoint, fmt.Sprintf(target.ResourceName, roomID)),
		httpmock.NewStringResponder(200, mockBody),
	)

	params := chatwork.RoomMembersUpdateParam{
		MembersAdminIDs:    param.IntArray(1, 2, 3),
		MembersMemberIDs:   optional.IntArray(10, 11, 12, 13),
		MembersReadonlyIDs: optional.IntArray(111, 112, 114),
	}

	actual, err := target.Update(roomID, params)
	assert.Nil(t, err)

	expected := rooms.RoomMembersAuthorityData{}
	err = json.Unmarshal([]byte(mockBody), &expected)

	assert.Nil(t, err)

	assert.Equal(t, expected, actual)
}
