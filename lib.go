package chatwork

import (
	"net/http"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/resource"
	"github.com/magicdrive/go-chatwork/api/resource/my"
	"github.com/magicdrive/go-chatwork/api/resource/rooms"
)

// ChatworkClientWrapper this chatwork library client type.
type ChatworkClientWrapper struct {
	Client *api.ChatworkAPIClient
}

// NewChatworkClient new this chatwork library client.
func NewChatworkClient(credential string) ChatworkClientWrapper {
	return ChatworkClientWrapper{
		Client: &api.ChatworkAPIClient{Credential: credential},
	}
}

// NewChatworkClientWithDetailed new this chatwork library client with detailed.
func NewChatworkClientWithDetailed(credential string, client *http.Client, altChatworkAPIHost string) ChatworkClientWrapper {
	return ChatworkClientWrapper{
		Client: api.NewChatworkAPIClient(credential, client, altChatworkAPIHost),
	}
}

// My chatwork api my resource.
func (c ChatworkClientWrapper) My() resource.MyResource {
	return resource.NewMyResource(c.Client)
}

// Me chatwork api me resource.
func (c ChatworkClientWrapper) Me() resource.MeResource {
	return resource.NewMeResource(c.Client)
}

// Contacts chatwork api contacts resource.
func (c ChatworkClientWrapper) Contacts() resource.ContactsResource {
	return resource.NewContactsResource(c.Client)
}

// IncomingRequests chatwork api incoming_requests resource.
func (c ChatworkClientWrapper) IncomingRequests() resource.IncomingRequestsResource {
	return resource.NewIncomingRequestsResource(c.Client)

}

// Rooms chatwork api rooms resource.
func (c ChatworkClientWrapper) Rooms() resource.RoomsResource {
	return resource.NewRoomsResource(c.Client)
}

// RoomUpdateParam chatwork api rooms update request param.
type RoomUpdateParam = resource.RoomsUpdateParam

// RoomCreateParam chatwork api rooms create request param.
type RoomCreateParam = resource.RoomsCreateParam

var (
	// IconPresetGroup    = resource.IconPresetGroup
	IconPresetGroup = resource.IconPresetGroup
	// IconPresetCheck    = resource.IconPresetCheck
	IconPresetCheck = resource.IconPresetCheck
	// IconPresetDocument = resource.IconPresetDocument
	IconPresetDocument = resource.IconPresetDocument
	// IconPresetMeeting  = resource.IconPresetMeeting
	IconPresetMeeting = resource.IconPresetMeeting
	// IconPresetEvent    = resource.IconPresetEvent
	IconPresetEvent = resource.IconPresetEvent
	// IconPresetProject  = resource.IconPresetProject
	IconPresetProject = resource.IconPresetProject
	// IconPresetBusiness = resource.IconPresetBusiness
	IconPresetBusiness = resource.IconPresetBusiness
	// IconPresetStudy    = resource.IconPresetStudy
	IconPresetStudy = resource.IconPresetStudy
	// IconPresetSecurity = resource.IconPresetSecurity
	IconPresetSecurity = resource.IconPresetSecurity
	// IconPresetStar     = resource.IconPresetStar
	IconPresetStar = resource.IconPresetStar
	// IconPresetIdea     = resource.IconPresetIdea
	IconPresetIdea = resource.IconPresetIdea
	// IconPresetHeart    = resource.IconPresetHeart
	IconPresetHeart = resource.IconPresetHeart
	// IconPresetMagcup   = resource.IconPresetMagcup
	IconPresetMagcup = resource.IconPresetMagcup
	// IconPresetBeer     = resource.IconPresetBeer
	IconPresetBeer = resource.IconPresetBeer
	// IconPresetMusic    = resource.IconPresetMusic
	IconPresetMusic = resource.IconPresetMusic
	// IconPresetSports   = resource.IconPresetSports
	IconPresetSports = resource.IconPresetSports
	// IconPresetTravel   = resource.IconPresetTravel
	IconPresetTravel = resource.IconPresetTravel
)

const (
	// RoomLeave  = resource.RoomLeave
	RoomLeave = resource.RoomLeave
	// RoomDelete = resource.RoomDelete
	RoomDelete = resource.RoomDelete
)

// MyTasksListParam chatwork api get my/task list request param.
type MyTasksListParam = my.MyTasksListParam

var (
	// MyTaskStatusDone = my.MyTaskStatusDone
	MyTaskStatusDone = my.MyTaskStatusDone
	// MyTaskStatusOpen = my.MyTaskStatusOpen
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

// RoomTaskPostParam chatwork api post rooms/task request param.
type RoomTaskPostParam = rooms.RoomTaskPostParam

var (
	// RoomTaskLimitTypeNone = rooms.RoomTaskLimitTypeNone
	RoomTaskLimitTypeNone = rooms.RoomTaskLimitTypeNone
	// RoomTaskLimitTypeDate = rooms.RoomTaskLimitTypeDate
	RoomTaskLimitTypeDate = rooms.RoomTaskLimitTypeDate
	// RoomTaskLimitTypeTime = rooms.RoomTaskLimitTypeTime
	RoomTaskLimitTypeTime = rooms.RoomTaskLimitTypeTime
)

var (
	// RoomTaskBodyStatusOpen = rooms.RoomTaskBodyStatusOpen
	RoomTaskBodyStatusOpen = rooms.RoomTaskBodyStatusOpen
	// RoomTaskBodyStatusDone = rooms.RoomTaskBodyStatusDone
	RoomTaskBodyStatusDone = rooms.RoomTaskBodyStatusDone
)
