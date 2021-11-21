package resource

import (
	"fmt"
	"net/http"

	json "github.com/goccy/go-json"
	"github.com/magicdrive/go-chatwork/api"
)

type IncommingRequestsResource struct {
	ResourceName string
	Credential   string
}

type IncommingRequestData struct {
	AccountID        int    `json:"account_id"`
	RoomID           int    `json:"room_id"`
	Name             string `json:"name"`
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int    `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string `json:"department"`
	AvatarImageURL   string `json:"avatar_image_url"`
}

func NewIncommingRequests(credential string) IncommingRequestsResource {
	data := IncommingRequestsResource{
		ResourceName: `/incomming_requests`,
		Credential:   credential,
	}
	return data
}

func (c IncommingRequestsResource) List() ([]IncommingRequestData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodGet,
		ResouceName: c.ResourceName,
		Params:      nil,
	}

	result := make([]IncommingRequestData, 0, 32)

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c IncommingRequestsResource) Accept(incomming_request_id int) (IncommingRequestData, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodPut,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, incomming_request_id),
		Params:      nil,
	}

	result := IncommingRequestData{}

	str, err := api.Call(spec)
	if err == nil {
		json.Unmarshal([]byte(str), &result)
	}

	return result, err
}

func (c IncommingRequestsResource) Delete(incomming_request_id int) (bool, error) {
	spec := api.ApiSpec{
		Credential:  c.Credential,
		Method:      http.MethodDelete,
		ResouceName: fmt.Sprintf("%s/%d", c.ResourceName, incomming_request_id),
		Params:      nil,
	}

	_, err := api.Call(spec)

	return err == nil, err
}
