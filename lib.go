package chatwork

import (
	"net/http"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/resource"
	"github.com/magicdrive/go-chatwork/api/resource/my"
	"github.com/magicdrive/go-chatwork/api/resource/rooms"
)

// ChatworkClient this chatwork library client type.
type ChatworkClient struct {
	Client *api.ChatworkApiClient
}

// NewChatworkClient new this chatwork library client.
func NewChatworkClient(credential string) ChatworkClient {
	return ChatworkClient{
		Client: &api.ChatworkApiClient{Credential: credential},
	}
}

// NewChatworkClientWithDetailed new this chatwork library client with detailed.
func NewChatworkClientWithDetailed(credential string, client *http.Client, alt_chatwork_api_host string) ChatworkClient {
	return ChatworkClient{
		Client: api.NewChatworkApiClient(credential, client, alt_chatwork_api_host),
	}
}

// My chatwork api my resource.
func (c ChatworkClient) My() resource.MyResource {
	return resource.NewMyResource(c.Client)
}

// Me chatwork api me resource.
func (c ChatworkClient) Me() resource.MeResource {
	return resource.NewMeResource(c.Client)
}

// Contacts chatwork api contacts resource.
func (c ChatworkClient) Contacts() resource.ContactsResource {
	return resource.NewContactsResource(c.Client)
}

// IncomingRequests chatwork api incoming_requests resource.
func (c ChatworkClient) IncomingRequests() resource.IncomingRequestsResource {
	return resource.NewIncomingRequestsResource(c.Client)

}

// Rooms chatwork api rooms resource.
func (c ChatworkClient) Rooms() resource.RoomsResource {
	return resource.NewRoomsResource(c.Client)
}

// RoomsUpdateParam chatwork api rooms update request param.
type RoomUpdateParam = resource.RoomsUpdateParam

// RoomsCreateParam chatwork api rooms create request param.
type RoomCreateParam = resource.RoomsCreateParam

var (
	IconPresetGroup    = resource.IconPresetGroup
	IconPresetCheck    = resource.IconPresetCheck
	IconPresetDocument = resource.IconPresetDocument
	IconPresetMeeting  = resource.IconPresetMeeting
	IconPresetEvent    = resource.IconPresetEvent
	IconPresetProject  = resource.IconPresetProject
	IconPresetBusiness = resource.IconPresetBusiness
	IconPresetStudy    = resource.IconPresetStudy
	IconPresetSecurity = resource.IconPresetSecurity
	IconPresetStar     = resource.IconPresetStar
	IconPresetIdea     = resource.IconPresetIdea
	IconPresetHeart    = resource.IconPresetHeart
	IconPresetMagcup   = resource.IconPresetMagcup
	IconPresetBeer     = resource.IconPresetBeer
	IconPresetMusic    = resource.IconPresetMusic
	IconPresetSports   = resource.IconPresetSports
	IconPresetTravel   = resource.IconPresetTravel
)

const (
	RoomLeave  = resource.RoomLeave
	RoomDelete = resource.RoomDelete
)

// MyTasksListParam chatwork api get my/task list request param.
type MyTasksListParam = my.MyTasksListParam

var (
	MyTaskStatusDone = my.MyTaskStatusDone
	MyTaskStatusOpen = my.MyTaskStatusOpen
)

// RoomLinkParam chatwork api get rooms/link list request param.
type RoomLinkParam = rooms.RoomLinkParam

// RoomMembersUpdateParam chatwork api update rooms/members update request param.
type RoomMembersUpdateParam = rooms.RoomMembersUpdateParam

// RoomMessagePostParam chatwork api post rooms/message request param.
type RoomMessagePostParam = rooms.RoomMessagePostParam

// RoomTasksListParam chatwork api get rooms/task list request param.
type RoomTasksListParam = rooms.RoomTasksListParam
// RoomTasksListParam chatwork api post rooms/task request param.
type RoomTaskPostParam = rooms.RoomTaskPostParam

var (
	RoomTaskLimitTypeNone = rooms.RoomTaskLimitTypeNone
	RoomTaskLimitTypeDate = rooms.RoomTaskLimitTypeDate
	RoomTaskLimitTypeTime = rooms.RoomTaskLimitTypeTime
)

var (
	RoomTaskBodyStatusOpen = rooms.RoomTaskBodyStatusOpen
	RoomTaskBodyStatusDone = rooms.RoomTaskBodyStatusDone
)
