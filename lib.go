package chatwork

import "github.com/magicdrive/go-chatwork/api/resource"

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
