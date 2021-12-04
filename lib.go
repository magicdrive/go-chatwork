package chatwork

import (
	"github.com/magicdrive/go-chatwork/api/resource"
	"github.com/magicdrive/go-chatwork/api/resource/my"
	"github.com/magicdrive/go-chatwork/api/resource/rooms"
)

type ChatworkClient struct {
	Credential string
}

func NewChatworkClient(credential string) ChatworkClient {
	return ChatworkClient{Credential: credential}
}

func (c ChatworkClient) My() resource.MyResource {
	return resource.NewMyResource(c.Credential)
}

func (c ChatworkClient) Me() resource.MeResource {
	return resource.NewMeResource(c.Credential)
}

func (c ChatworkClient) Contacts() resource.ContactsResource {
	return resource.NewContactsResource(c.Credential)
}

func (c ChatworkClient) IncomingRequests() resource.IncomingRequestsResource {
	return resource.NewIncomingRequestsResource(c.Credential)

}
func (c ChatworkClient) Rooms() resource.RoomsResource {
	return resource.NewRoomsResource(c.Credential)
}

type RoomUpdatePram = resource.RoomsUpdateParam
type RoomCreateData = resource.RoomsCreateParam

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
