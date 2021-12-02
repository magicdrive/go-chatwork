package resource

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
)

type IncomingRequestsResource struct {
	ResourceName string
	Credential   string
}

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

func NewIncomingRequestsResource(credential string) IncomingRequestsResource {
	data := IncomingRequestsResource{
		ResourceName: `/incoming_requests`,
		Credential:   credential,
	}
	return data
}

func (c IncomingRequestsResource) List() ([]IncomingRequestData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := make([]IncomingRequestData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c IncomingRequestsResource) Accept(incoming_request_id int) (IncomingRequestData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, incoming_request_id),
		Params:      nil,
	}

	result := IncomingRequestData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c IncomingRequestsResource) Delete(incoming_request_id int) error {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, incoming_request_id),
		Params:      nil,
	}

	_, err := api.Call(spec)

	return err
}
