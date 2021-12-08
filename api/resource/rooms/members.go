package rooms

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

// RoomMembersResource chatwork api rooms/members resouce.
type RoomMembersResource struct {
	ResourceName string
	Client       *api.ChatworkAPIClient
}

// RoomMemberData chatwork api resp rooms/member data.
type RoomMemberData struct {
	AccountID        int    `json:"account_id"`
	Role             string `json:"role"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageURL   string `json:"avatar_image_url"`
}

// RoomMembersAuthorityData chatwork api resp rooms/member autority data.
type RoomMembersAuthorityData struct {
	Admin    []int `json:"admin"`
	Member   []int `json:"member"`
	Readonly []int `json:"readonly"`
}

// RoomMembersUpdateParam chatwork api update rooms/members request param.
type RoomMembersUpdateParam struct {
	MembersAdminIDs    *param.IntArrayParam       `json:"members_admin_ids"`
	MembersMemberIDs   *optional.NullableIntArray `json:"members_member_ids"`
	MembersReadonlyIDs *optional.NullableIntArray `json:"members_readonly_ids"`
}

// NewRoomMembersResource new chatwork api rooms/members resource.
func NewRoomMembersResource(parent string, client *api.ChatworkAPIClient) RoomMembersResource {
	data := RoomMembersResource{
		ResourceName: parent + `/%d/members`,
		Client:   client,
	}
	return data
}

// List chatwork api get rooms/members list.
func (c RoomMembersResource) List(roomID int, force *optional.NullableBool) ([]RoomMemberData, error) {

	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: fmt.Sprintf(c.ResourceName, roomID),
		Params: map[string]*optional.NullableString{
			"force": force.ToNullableString(),
		},
	}

	result := make([]RoomMemberData, 0, 32)

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Update chatwork api update rooms/members.
func (c RoomMembersResource) Update(roomID int, params RoomMembersUpdateParam) (RoomMembersAuthorityData, error) {

	b, _ := json.Marshal(params)
	p, _ := api.JSONToMap(b)

	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPost,
		ResouceName: fmt.Sprintf(c.ResourceName, roomID),
		Params:      p,
	}

	result := RoomMembersAuthorityData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
