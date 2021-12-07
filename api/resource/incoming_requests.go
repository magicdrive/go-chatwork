package resource

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
)

// IncomingRequestsResource chatwork api incoming_requests resource.
type IncomingRequestsResource struct {
	ResourceName string
	Client       *api.ChatworkApiClient
}

// IncomingRequestData chatwork api resp incoming_request data.
type IncomingRequestData struct {
	AccountID        int    `json:"account_id"`
	RoomID           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageURL   string `json:"avatar_image_url"`
}

// NewIncomingRequestsResource chatwork api new incoming_request resource.
func NewIncomingRequestsResource(client *api.ChatworkApiClient) IncomingRequestsResource {
	data := IncomingRequestsResource{
		ResourceName: `/incoming_requests`,
		Client:       client,
	}
	return data
}

// List chatwork api get incoming_request list.
func (c IncomingRequestsResource) List() ([]IncomingRequestData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := make([]IncomingRequestData, 0, 32)

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Accept chatwork api accept incoming_request.
func (c IncomingRequestsResource) Accept(incoming_request_id int) (IncomingRequestData, error) {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, incoming_request_id),
		Params:      nil,
	}

	result := IncomingRequestData{}

	str, err := c.Client.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

// Delete chatwork api delete incoming_request.
func (c IncomingRequestsResource) Delete(incoming_request_id int) error {
	spec := api.ApiSpec{
		Credential:  c.Client.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, incoming_request_id),
		Params:      nil,
	}

	_, err := c.Client.Call(spec)

	return err
}
