package resource

import (
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
)

// MeResource chatwork api me resource.
type MeResource struct {
	ResourceName string
	Client       *api.ChatworkAPIClient
}

// MeData chatwork api me resp data.
type MeData struct {
	AccountID        int    `json:"account_id"`
	RoomID           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	Title            string `json:"title"`
	URL              string `json:"url"`
	Introduction     string `json:"introduction"`
	Mail             string `json:"mail"`
	TelOrganization  string `json:"tel_organization"`
	TelExtension     string `json:"tel_extension"`
	TelMobile        string `json:"tel_mobile"`
	Skype            string `json:"skype"`
	Facebook         string `json:"facebook"`
	Twitter          string `json:"twitter"`
	AvatarImageURL   string `json:"avatar_image_url"`
	LoginMail        string `json:"login_mail"`
}

// NewMeResource new chatwork api me resource.
func NewMeResource(client *api.ChatworkAPIClient) MeResource {
	data := MeResource{
		ResourceName: `/me`,
		Client:       client,
	}
	return data
}

// Get new chatwork api get me data.
func (c MeResource) Get() (MeData, error) {
	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := MeData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
