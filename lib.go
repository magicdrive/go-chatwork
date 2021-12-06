package chatwork

import (
	"net/http"

	"github.com/magicdrive/go-chatwork/api"
	"github.com/magicdrive/go-chatwork/api/resource"
	"github.com/magicdrive/go-chatwork/api/resource/my"
	"github.com/magicdrive/go-chatwork/api/resource/rooms"
)

type ChatworkClient struct {
	Client *api.ChatworkApiClient
}

func NewChatworkClient(credential string) ChatworkClient {
	return ChatworkClient{
		Client: &api.ChatworkApiClient{Credential: credential},
	}
}

func NewChatworkClientWithDetailed(credential string, client *http.Client, alt_chatwork_api_host string) ChatworkClient {
	return ChatworkClient{
		Client: api.NewChatworkApiClient(credential, client, alt_chatwork_api_host),
	}
}

func (c ChatworkClient) My() resource.MyResource {
	return resource.NewMyResource(c.Client)
}

func (c ChatworkClient) Me() resource.MeResource {
	return resource.NewMeResource(c.Client)
}

func (c ChatworkClient) Contacts() resource.ContactsResource {
	return resource.NewContactsResource(c.Client)
}

func (c ChatworkClient) IncomingRequests() resource.IncomingRequestsResource {
	return resource.NewIncomingRequestsResource(c.Client)

}
func (c ChatworkClient) Rooms() resource.RoomsResource {
	return resource.NewRoomsResource(c.Client)
}

type RoomUpdateParam = resource.RoomsUpdateParam
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

type MyTasksListParam = my.MyTasksListParam

var (
	MyTaskStatusDone = my.MyTaskStatusDone
	MyTaskStatusOpen = my.MyTaskStatusOpen
)

type RoomLinkParam = rooms.RoomLinkParam

type RoomMembersUpdateParam = rooms.RoomMembersUpdateParam

type RoomMessagePostParam = rooms.RoomMessagePostParam

type RoomTasksListParam = rooms.RoomTasksListParam
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
