package resource

import (
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
)

// ContactsResource chatwork api contacts resource.
type ContactsResource struct {
	ResourceName string
	Client       *api.ChatworkAPIClient
}

// ContactData chatwork api resp contact data.
type ContactData struct {
	AccountID        int    `json:"account_id"`
	RoomID           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageURL   string `json:"avatar_image_url"`
}

// NewContactsResource new chatwork api contacts resource.
func NewContactsResource(client *api.ChatworkAPIClient) ContactsResource {
	data := ContactsResource{
		ResourceName: `/contacts`,
		Client:       client,
	}
	return data
}

// List chatwork api resp contact data.
func (c ContactsResource) List() ([]ContactData, error) {
	spec := api.APISpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := make([]ContactData, 0, 32)

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}
