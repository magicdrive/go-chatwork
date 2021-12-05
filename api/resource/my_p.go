package resource

import (
	"github.com/magicdrive/go-chatwork/api"
	my "github.com/magicdrive/go-chatwork/api/resource/my"
)

type MyResource struct {
	ResourceName string
	Client       *api.ChatworkApiClient
}

func NewMyResource(client *api.ChatworkApiClient) MyResource {
	data := MyResource{
		ResourceName: `/my`,
		Client:       client,
	}
	return data
}

func (c MyResource) Status() my.MyStatusResource {
	return my.NewMyStatus(c.ResourceName, c.Client)
}

func (c MyResource) Tasks() my.MyTasksResource {
	return my.NewMyTasks(c.ResourceName, c.Client)
}
